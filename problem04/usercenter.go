package problem04

import (
	"context"
	"errors"
)

// 简易代码示例

type UserInfo struct {
	UserID int
	Phone string
	Password string
	Salt string
	Token string
}

func NewUserInfo(phone, rawPassword string) *UserInfo {
	salt := createSalt()
	password := encryptPassword(rawPassword, salt)
	return &UserInfo{
		Password: password,
		Salt: salt,
		Phone: phone,
	}
}

func (user *UserInfo) verify(phone, rawPassword string) bool {
	if user.Phone != phone {
		return false
	}
	if user.Password != encryptPassword(rawPassword, user.Salt) {
		return false
	}
	return true
}

func generateToken() string {
	return ""
}

func createSalt() string {
	return ""
}
func encryptPassword(password string, salt string) (encrypted string) {
	// TODO 密码加密
	_, _ = password, salt
	return ""
}


type UserCenter interface {
	Login(ctx context.Context, phone string, password string) (*UserInfo, error)
	Register(ctx context.Context, phone string, password string) (*UserInfo, error)
	RebindPhone(ctx context.Context, token string, newPhone string) error
}

type RedisUserCenter struct {
	idGenerator IDGenerator
}

func (r *RedisUserCenter) Login(ctx context.Context, phone string, password string) (*UserInfo, error) {
	userID, find := r.getUserIDByPhone(ctx, phone)
	if !find {
		return nil, errors.New("手机号或密码错误")
	}

	userInfo, find := r.getUserInfoByID(ctx, userID)

	if !find || !userInfo.verify(phone, password) {
		return nil, errors.New("手机号或密码错误")
	}
	userInfo.Token = generateToken()
	r.saveToken(ctx, userInfo.Token, userID)
	return userInfo, nil
}

func (r *RedisUserCenter) Register(ctx context.Context, phone string, password string) (*UserInfo, error) {
	userID, find := r.getUserIDByPhone(ctx, phone)
	if find {
		userInfo, find := r.getUserInfoByID(ctx, userID)
		if find && userInfo.Phone == phone {
			return nil, errors.New("手机号已注册")
		}
	}

	info := NewUserInfo(phone, password)

	for i:=0; i<3; i++{
		userID := int(r.idGenerator.GenerateID())
		r.savePhone(ctx, phone, userID)
		r.saveToken(ctx, info.Token, userID)

		info.UserID = userID
		if r.createUserInfo(ctx, info) {
			return info, nil
		}
	}
	return nil, errors.New("服务器内部错误")
}

func (r *RedisUserCenter) RebindPhone(ctx context.Context, token string, newPhone string) error {
	_, _, _ = ctx, token, newPhone
	return nil
}

func (r *RedisUserCenter) createUserInfo(ctx context.Context, info *UserInfo) bool {
	// TODO 将 userInfo 对象存入 redis 中, 如果该 key 已经存在, 直接返回 false, 存入成功返回 true
	_, _ = ctx, info
	return true
}

func (r *RedisUserCenter) getUserInfoByID(ctx context.Context, userID int) (*UserInfo, bool) {
	// TODO 从 redis 中读取用户对象
	_, _ = ctx, userID
	return nil, false
}

func (r *RedisUserCenter) getUserIDByPhone(ctx context.Context, phone string) (int, bool) {
	// TODO 从过 phone 获取 userID
	// HGET phones 13767556733
	_, _ = ctx, phone
	return 0, false
}

func (r *RedisUserCenter) savePhone(ctx context.Context, phone string, userID int) {
	// TODO 将 phone-userID 键值对存入到 phones Hash 中
	_, _, _ = ctx, phone, userID
}

func (r *RedisUserCenter) saveToken(ctx context.Context, token string, userID int) {
	// TODO 将 token-userID 键值对存入到 tokens Hash 中
	_, _, _ = ctx, token, userID
}

