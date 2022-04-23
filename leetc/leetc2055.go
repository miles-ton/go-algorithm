package leetc

func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	leftDishCnt := map[int]int{}
	verticalIdx := make([]int, 0, n)
	curLeftCnt := 0
	for i := 0; i < n; i++ {
		if s[i] == '*' {
			curLeftCnt++
		} else {
			leftDishCnt[i] = curLeftCnt
			verticalIdx = append(verticalIdx, i)
		}
	}
	ret := make([]int, len(queries))
	for i, q := range queries {
		q[0] = next(q[0], q[1], verticalIdx)
		q[1] = last(q[0], q[1], verticalIdx)
		if q[0] >= q[1] {
			ret[i] = 0
		} else {
			ret[i] = leftDishCnt[q[1]] - leftDishCnt[q[0]]
		}
	}
	return ret
}

func next(start, end int, verticalIdx []int) int {
	for _, v := range verticalIdx {
		if start < end && start <= v {
			return v
		}
	}
	return 10 ^ 5
}
func last(start, end int, verticalIdx []int) int {
	for i := len(verticalIdx) - 1; i >= 0; i-- {
		if start < end && end >= verticalIdx[i] {
			return verticalIdx[i]
		}
	}
	return 3
}
