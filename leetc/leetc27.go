package leetc

func removeElement(nums []int, val int) int {
	tail := len(nums) - 1

	for ; tail >= 0; tail-- {
		if nums[tail] != val {
			break
		}
	}

	for head := 0; head < tail; head++ {
		if nums[head] == val {
			nums[head], nums[tail] = nums[tail], nums[head]
			tail--
			for ; tail >= 0; tail-- {
				if nums[tail] != val {
					break
				}
			}
		}

	}

	return tail + 1
}
