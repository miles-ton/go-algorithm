package leetc

func SearchRange(nums []int, target int) []int {
	ret := []int{-1, -1}
	if len(nums) == 0 {
		return ret
	}
	s, m, e := binarySearch(nums, 0, len(nums)-1, target)
	if m == -1 {
		return ret
	}

	realM := m
	ret = []int{m, m}
	for s, m, _ = binarySearch(nums, s, m-1, target); m != -1; s, m, _ = binarySearch(nums, s, m-1, target) {
		ret[0] = m
	}
	m = realM
	for _, m, e = binarySearch(nums, m+1, e, target); m != -1; _, m, e = binarySearch(nums, m+1, e, target) {
		ret[1] = m
	}

	return ret
}

func binarySearch(nums []int, start, end, target int) (s, m, e int) {
	if start > end {
		return -1, -1, -1
	}

	mid := (start + end) / 2
	actual := nums[mid]
	if target == actual {
		return start, mid, end
	} else if target < actual {
		return binarySearch(nums, start, mid-1, target)
	} else {
		return binarySearch(nums, mid+1, end, target)
	}
}
