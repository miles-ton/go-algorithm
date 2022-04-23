package leetc

import "sort"

func SubsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	var storedElem []int
	var recur func([]int, []int)
	recur = func(elem, nums []int) {
		if len(elem) == 0 {
			e := make([]int, len(storedElem))
			copy(e, storedElem)
			res = append(res, e)
			return
		}
		if len(elem) > len(nums) {
			return
		}
		last := 11
		for i := 0; i < len(nums); i++ {
			if nums[i] == last {
				continue
			}
			elem[0] = nums[i]
			last = nums[i]
			recur(elem[1:], nums[i+1:])
		}
	}
	res = append(res, []int{})
	for i := 0; i < len(nums); i++ {
		storedElem = make([]int, i+1)
		recur(storedElem, nums)
	}

	return res
}
