package leetc

func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	dpSeq, dpInver := make([]int, n), make([]int, n)
	dpSeq[0] = time
	for i := 1; i < n; i++ {
		if security[i] <= security[i-1] {
			dpSeq[i] = dpSeq[i-1] - 1
		} else {
			dpSeq[i] = time
		}
	}
	dpInver[n-1] = time
	for i := n - 2; i >= 0; i-- {
		if security[i] <= security[i+1] {
			dpInver[i] = dpInver[i+1] - 1
		} else {
			dpInver[i] = time
		}
	}
	ret := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if dpSeq[i] <= 0 && dpInver[i] <= 0 {
			ret = append(ret, i)
		}
	}

	return ret
}
