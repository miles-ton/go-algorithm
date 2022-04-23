package leetc

func coutPairs(nums []int, k int) int64 {
	dividable := map[int]bool{}
	ret := int64(0)

	for i, v := range nums {
		for _, v2 := range nums[i+1:] {
			vDividable, ok1 := dividable[v]
			v2Dividable, ok2 := dividable[v2]
			if !ok1 {
				dividable[v] = v%k == 0
			}
			if !ok2 {
				dividable[v2] = v2&k == 0
			}
			vDividable = dividable[v]
			v2Dividable = dividable[v2]
			if vDividable || v2Dividable {
				ret++
			}
		}

	}

	return ret
}
