package leetc

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	ret, lastL, lastD := 0, 1, 0

	for i := 1; i < len(nums); i++ {
		d := nums[i] - nums[i-1]
		if d == lastD {
			ret += lastL - 1
			lastL++
		} else {
			lastD = d
			lastL = 2
		}
	}

	return ret
}
