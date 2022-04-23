package leetc

func canJump(nums []int) bool {
	maxxx := nums[0]
	for cur := 0; cur <= maxxx && maxxx < len(nums); cur++ {
		maxxx = max55(maxxx, nums[cur]+cur)
	}
	return maxxx >= len(nums)-1
}

func max55(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
