package leetc

func minSubArrayLen(target int, nums []int) int {
	head, ret, sum := 0, 100000+1, 0
	for i, v := range nums {
		sum += v
		if sum >= target {
			ret = min209(ret, i-head+1)
			for head < i {
				sum -= nums[head]
				head++
				if sum >= target {
					ret = min209(ret, i-head+1)
				} else {
					break
				}
			}
		}
	}
	if ret == 100001 {
		ret = 0
	}
	return ret
}

func min209(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
