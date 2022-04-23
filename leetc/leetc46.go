package leetc

func permute(nums []int) [][]int {
	n := len(nums)
	var res [][]int
	var recur func(nums []int, alloc int)
	recur = func(nums []int, alloc int) {
		if alloc == n {
			res = append(res, append(make([]int, 0, n), nums...))
			return
		}
		for i := alloc; i < n; i++ {
			nums[alloc], nums[i] = nums[i], nums[alloc]
			recur(nums, alloc+1)
			nums[alloc], nums[i] = nums[i], nums[alloc]
		}

	}

	recur(nums, 0)

	return res
}
