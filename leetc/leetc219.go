package leetc

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {
	vi := make(map[int]int)

	for i, v := range nums {
		mv, ok := vi[v]
		if ok {
			if int(math.Abs(float64(i-mv))) <= k {
				return true
			} else {
				vi[v] = i
			}
		} else {
			vi[v] = i
		}
	}
	return false
}
