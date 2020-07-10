package problem04

import (
	"context"
	"errors"
)

// 简易代码示例

type UserInfo struct {
	UserID int32
	Phone string
	Password string
	Salt string
	Token string
}

func NewUserInfo() *UserInfo {
	return nil
}


type UserCenter interface {
	Login(ctx context.Context, phone string, password string) (*UserInfo, error)
	Register(ctx context.Context, phone string, password string) (*UserInfo, error)
}

type RedisUserCenter struct {
	idGenerator IDGenerator
}

func (r *RedisUserCenter) Login(ctx context.Context, phone string, password string) (*UserInfo, error) {
	userID, find := r.getUserIDByPhone(phone)
	if !find {
		return nil, errors.New("手机号或密码错误")
	}

	userInfo, find := r.getUserInfoByID(userID)
	if !find || userInfo.Phone != phone {
		return nil, errors.New("手机号或密码错误")
	}
	token := r.generateToken()
	r.saveToken(phone, token)
	return nil, errors.New("服务器内部错误")
}

func (r *RedisUserCenter) Register(ctx context.Context, phone string, password string) (*UserInfo, error) {
	userID, find := r.getUserIDByPhone(phone)
	if find {
		userInfo, find := r.getUserInfoByID(userID)
		if find && userInfo.Phone == phone {
			return nil, errors.New("手机号已注册")
		}
	}

	info := NewUserInfo()
	for i:=0; i<3; i++{
		userID := int(r.idGenerator.GenerateID())
		r.savePhone(phone, userID)
		r.saveToken(phone, info.Token)

		if r.CreateUserInfo(info) {
			return info, nil
		}
	}
	return nil, errors.New("服务器内部错误")
}

func (r *RedisUserCenter) CreateUserInfo(info *UserInfo) bool {
	return true
}

func (r *RedisUserCenter) getUserInfoByID(userID int) (*UserInfo, bool) {
	return nil, false
}
func (r *RedisUserCenter) generateToken() string {
	return ""
}
func (r *RedisUserCenter) encryptPassword(password string) (encrypted, salt string) {
	return "", ""
}

func (r *RedisUserCenter) getUserIDByPhone(phone string) (int, bool) {
	// HGET phones 13767556733
	return 0, false
}

func (r *RedisUserCenter) savePhone(phone string, userID int) {

}

func (r *RedisUserCenter) saveToken(phone string, token string) {

}

