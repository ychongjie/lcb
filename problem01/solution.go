package problem01

// 现有整型数组 a 、整型数组 b、以及整型v。请使⽤指定编程语⾔编写函数，
// 判断是否可以从 a 中选择⼀个数，b 中选择⼀个数，⼆者相加等于 v，如可以返回 true ，否则返回 false
func TargetPairExists(arr1, arr2 []int, targetSum int) bool {
	if len(arr1) == 0 || len(arr2) == 0 {
		return false
	}
	set := make(map[int]struct{}, len(arr1))
	for _, num := range arr1 {
		set[num] = struct{}{}
	}

	for _, num := range arr2 {
		_, ok := set[targetSum-num]
		if ok {
			return true
		}
	}

	return false
}