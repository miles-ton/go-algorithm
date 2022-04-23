package leetc

func minOperations(nums []int) int {
	ret := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			continue
		}
		incr := nums[i-1] - nums[i] + 1
		nums[i] += incr
		ret += incr
	}
	return ret
}
