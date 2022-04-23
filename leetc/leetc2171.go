package leetc

import "sort"

func MinimumRemoval(beans []int) (ret int64) {
	n := len(beans)
	sort.Ints(beans)
	for _, v := range beans {
		ret += int64(v - beans[0])
	}
	cur := ret
	for i, v := range beans[:n-1] {
		cur = cur + int64(v-(beans[i+1]-v)*(n-1-i))
		ret = min2171(ret, cur)
	}
	return
}

func min2171(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
