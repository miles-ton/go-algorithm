package leetc

func sortColors(nums []int) {
	oneI, twoI := 0, len(nums)-1
	for i := 0; i <= twoI; {
		if nums[i] == 0 {
			copy(nums[oneI+1:i+1], nums[oneI:i])
			nums[oneI] = 0
			oneI++
			if oneI > i {
				i = oneI
			}
		} else if nums[i] == 2 {
			copy(nums[i:twoI], nums[i+1:twoI+1])
			nums[twoI] = 2
			twoI--
		} else {
			i++
		}
	}
}
