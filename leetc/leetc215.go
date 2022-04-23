package leetc

import "sort"

func findKthLargest(nums []int, k int) int {

	sort.Ints(nums)

	return nums[len(nums)-1-k]
}
