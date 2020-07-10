package problem02

import (
	"errors"
	"math"
	"strconv"
)

//现要求游戏玩家的⽤户名是 user 后加上不重复且随机的正整数，

//请参考 PostgreSQL 中 pseudo_encrypt() 的原理，使⽤指定编程语⾔，实现⼀个⾼效的⽤户名⽣成算法。

const prefix = "user_"

// 输入参数为递增的正整数, 例如可以为数据库递增主键
func generateUsername(index int) (username string, err error) {
	if index < 0 {
		return "", errors.New("index should be positive integer")
	}

	return prefix + strconv.Itoa(int(pseudoEncrypt(int32(index)))), nil
}


//https://wiki.postgresql.org/wiki/Pseudo_encrypt
func pseudoEncrypt(value int32) int32 {
	var l1, r1, l2, r2 int32
	l1 = (value >> 16) & 0xffff
	r1 = value & 0xffff
	for i:=0; i<3; i++ {
		/* round() is used to produce the same values as the
		   plpgsql implementation that does an SQL cast to INT */
		l2, r2 = r1, l1 ^ int32(math.Round((float64((1366*r1 + 150889) % 714025) / 714025.0) * 32767))
		l1, r1 = l2, r2
	}
	return (r1 << 16) + l1
}

