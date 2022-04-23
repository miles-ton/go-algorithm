package leetc

func MinimumOperations(nums []int) int {
	n := len(nums)
	var one, two []int
	for i := 0; i < n; i += 2 {
		one = append(one, nums[i])
	}
	for i := 1; i < n; i += 2 {
		two = append(two, nums[i])
	}
	oneN := maxAndSecSubSeq(one)
	twoN := maxAndSecSubSeq(two)
	if oneN[0] != twoN[0] {
		return n - (oneN[1] + twoN[1])
	} else {
		if oneN[1]+twoN[3] > oneN[3]+twoN[1] {
			return n - (oneN[1] + twoN[3])
		} else {
			return n - (oneN[3] + twoN[1])
		}
	}
}

func maxAndSecSubSeq(seq []int) (ret []int) {
	cnt := map[int]int{}
	for _, v := range seq {
		if _, has := cnt[v]; has {
			cnt[v]++
		} else {
			cnt[v] = 1
		}
	}
	ret = make([]int, 4)
	for k, v := range cnt {
		if v > ret[3] {
			ret[2], ret[3] = k, v
		}
		if v >= ret[1] {
			a, b := ret[0], ret[1]
			ret[0], ret[1] = ret[2], ret[3]
			ret[2], ret[3] = a, b
		}
	}

	return
}
