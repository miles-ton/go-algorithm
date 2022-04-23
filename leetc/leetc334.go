package leetc

func increasingTriplet(nums []int) bool {
	one, two := 0, -1

	for i, v := range nums {
		if two != -1 && v > nums[two] {
			return true
		} else if v > nums[one] {
			two = i
		} else {
			one = i
		}
	}
	return false
}
