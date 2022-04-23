package leetc

func getRow(rowIndex int) []int {
	ret := make([]int, rowIndex+1)
	for i := 0; i < rowIndex+1; i++ {
		ret[0] = 1
		ret[i] = 1
		for j := i - 1; j > 0; j-- {
			ret[j] = ret[j-1] + ret[j]
		}
	}

	return ret
}
