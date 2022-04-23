package leetc

func subsets(nums []int) [][]int {
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
		for i := 0; i < len(nums); i++ {
			elem[0] = nums[i]
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
