package leetc

func search(nums []int, target int) int {
	return recur704(0, len(nums)-1, nums, target)
}

func recur704(left, right int, nums []int, target int) int {
	if left > right {
		return -1
	}
	m := (left + right) / 2
	if nums[m] == target {
		return m
	} else if target < nums[m] {
		return recur704(left, m-1, nums, target)
	} else {
		return recur704(m+1, right, nums, target)
	}
}
