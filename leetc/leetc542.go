package leetc

func singleNonDuplicate(nums []int) int {
	return binary(nums)
}

func binary(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	mid := len(nums) / 2
	if nums[mid] != nums[mid+1] && nums[mid] != nums[mid-1] {
		return nums[mid]
	}
	sameAsleft := false
	if mid%2 == 1 {
		sameAsleft = true
	}
	if sameAsleft {
		if nums[mid] == nums[mid-1] {
			return binary(nums[mid+1:])
		}
		return binary(nums[:mid])
	}
	if nums[mid] == nums[mid+1] {
		return binary(nums[mid:])
	}
	return binary(nums[:mid-1])
}
