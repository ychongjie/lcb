package problem03

import (
	"math"
	"math/rand"
	"time"
)

// 如果要求⽤户名均为8位，并且即便知道所⽤算法，也⽆法预测⽤户名序列，应该如何改进算法

// 假设输入的 value 序列是从 0 开始严格递增的, 该函数需满足:

const maxValue = int32(99999999)

// 1. 返回值在 [0, 99999999] 区间内, 参考 https://wiki.postgresql.org/wiki/Pseudo_encrypt_constrained_to_an_arbitrary_range
func boundedPseudoEncrypt(value int32) int32 {
	for {
		value = pseudoEncrypt(value)
		if value <= maxValue {
			return value
		}
	}
}

func pseudoEncrypt(value int32) int32 {
	var l1, r1, l2, r2 int32
	l1 = (value >> 13) & (8192 - 1)
	r1 = value & (8192 - 1)
	for i:=0; i<3; i++ {
		/* round() is used to produce the same values as the
		   plpgsql implementation that does an SQL cast to INT */
		l2, r2 = r1, l1 ^ int32(math.Round((float64((1366*r1 + 150889) % 714025) / 714025.0) * (8192 - 1)))
		l1, r1 = l2, r2
	}
	return (l1 << 13) + r1
}

// 2 无法预测用户名序列, 也就是说要在输入输出之间增加一定的随机性
// 想到的第一种办法是将生成值的低位和一个随机值做异或
// 但这种方法会极大的增加值碰撞的概率, 不太可取
func randomBoundedPseudoEncrypt(value int32) int32 {
	encryptValue := boundedPseudoEncrypt(value)
	rand.Seed(time.Now().UnixNano())
	r := int32(rand.Intn(8))
	return encryptValue ^ r
}

// 另一种方法是让核心函数的参数运行时随机生成
var a, b, c int
func init() {
	rand.Seed(time.Now().UnixNano())
	a = 1000 + rand.Intn(500)
	b = 150000 + rand.Intn(1000)
	c = 700000 + rand.Intn(1000)
}

func boundedPseudoEncrypt2(value int32) int32 {
	for {
		value = pseudoEncrypt2(value)
		if value <= maxValue {
			return value
		}
	}
}

func pseudoEncrypt2(value int32) int32 {
	var l1, r1, l2, r2 int32
	l1 = (value >> 13) & (8192 - 1)
	r1 = value & (8192 - 1)
	for i:=0; i<3; i++ {
		/* round() is used to produce the same values as the
		   plpgsql implementation that does an SQL cast to INT */
		l2, r2 = r1, l1 ^ int32(math.Round((float64((a*int(r1) + b) % c) / float64(c)) * (8192 - 1)))
		l1, r1 = l2, r2
	}
	return (l1 << 13) + r1
}