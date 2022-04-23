package leetc

func dominantIndex(nums []int) int {

	first, second, ret := -1, -1, -1
	for i, v := range nums {
		if v > second {
			first = second
			second = v
			ret = i
		} else if v > first {
			first = v
		}
	}

	if second >= first<<1 {
		return ret
	} else {
		return -1
	}
}
