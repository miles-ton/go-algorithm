package leetc

func moveZeroes(nums []int) {
	head := 0
	for i, v := range nums {
		if v != 0 {
			nums[head], nums[i] = nums[i], nums[head]
			head++
		}
	}
}
