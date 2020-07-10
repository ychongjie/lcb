package problem03

import (
	"math/rand"
	"testing"
)

// 验证一下该算法出现重复的概率是否够低
// 一千万个数中没有出现一次重复, 碰撞概率可以接受
func Test_boundedPseudoEncrypt_noRepeat(t *testing.T) {
	dict := make(map[int32]struct{})
	for i:=0; i<9999999; i++ {
		v := boundedPseudoEncrypt(int32(i))
		if _, ok := dict[v]; ok {
			t.Errorf("found repeat value for %d", i)
		}
		dict[v] = struct{}{}
	}
}

func Test_boundedPseudoEncrypt2_noRepeat(t *testing.T) {
	dict := make(map[int32]struct{})
	for i:=0; i<9999999; i++ {
		v := boundedPseudoEncrypt2(int32(i))
		if _, ok := dict[v]; ok {
			t.Errorf("found repeat value for %d", i)
		}
		dict[v] = struct{}{}
	}
}

// 通过异或一个随机数来让生成序列不可预测, 效果并不好
// 因为算法的碰撞概率太高, 比直接用 rand 没有很大的优势
func Test_randomBoundedPseudoEncrypt_repeatCounts(t *testing.T) {
	conflictCount1 := countCollision(1000000, randomBoundedPseudoEncrypt)
	t.Logf("conflict count for randomBoundedPseudoEncrypt: %d", conflictCount1)

	conflictCount2 := countCollision(1000000, func(i int32) int32 {
		return int32(rand.Intn(int(maxValue + 1)))
	})
	t.Logf("conflict count for rand: %d", conflictCount2)
}

func countCollision (n int, cal func(int32) int32) int {
	dict := make(map[int32]struct{})
	collision := 0
	for i:=0; i<n; i++ {
		v := cal(int32(i))
		if _, ok := dict[v]; ok {
			collision++
		}
		dict[v] = struct{}{}
	}
	return collision
}