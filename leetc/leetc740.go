package leetc

func deleteAndEarn(nums []int) int {
	var maxVal int
	for _, v := range nums {
		maxVal = max(maxVal, v)
	}
	sum := make([]int, maxVal+1)
	for _, v := range nums {
		sum[v] += v
	}
	return rob(sum)
}

func rob(nums []int) int {
	r := nums[0]
	nr := 0
	for _, v := range nums {
		tmp := r
		r = nr + v
		nr = max(nr, tmp)
	}
	return max(nr, r)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
