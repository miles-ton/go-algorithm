package leetc

import "math"

func Jump(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 0

	for i := 1; i < len(nums); i++ {
		step := math.MaxInt32
		for j := 0; j < i; j++ {
			if i-j <= nums[j] {
				step = min45(step, dp[j]+1)
			}
		}
		dp[i] = step
	}
	return dp[len(nums)-1]
}

func min45(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
