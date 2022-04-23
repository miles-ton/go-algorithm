package leetc

func SortArray(nums []int) []int {
	recur912(nums, 0, len(nums)-1)
	return nums
}

func recur912(nums []int, l, r int) {
	if l >= r {
		return
	}
	left, right := l, r

	for left < right {
		for left < right && nums[left] <= nums[right] {
			left++
		}
		swap(nums, left, right)
		for left < right && nums[right] > nums[left] {
			right--
		}
		swap(nums, left, right)
	}
	recur912(nums, l, left-1)
	recur912(nums, right+1, r)
}

func swap(nums []int, a, b int) {
	tmp := nums[a]
	nums[a] = nums[b]
	nums[b] = tmp
}
