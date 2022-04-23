package leetc

import "sort"

func LargestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	maxLen, maxIdx := 1, -1
	dp := make([]int, len(nums))
	prev := make([]int, len(nums))
	for i := range nums {
		dp[i] = 1
		prev[i] = -1
	}

	for i, v := range nums {
		for j := i - 1; j >= 0; j-- {
			if v%nums[j] == 0 {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					prev[i] = j
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
			maxIdx = i
		}
	}
	return getSeq(nums, prev, maxIdx)
}

func getSeq(nums, prev []int, start int) []int {
	if start == -1 {
		return nums[0:1]
	}
	sub := make([]int, 0, len(nums))
	for ; start != -1; start = prev[start] {
		sub = append(sub, nums[start])
	}
	return sub[:]
}
