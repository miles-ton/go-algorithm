package leetc

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)
	ret[0] = make([]int, 1)
	ret[0][0] = 1
	for i := 1; i < numRows; i++ {
		ret[i] = make([]int, i+1)
		ret[i][0] = 1
		ret[i][i] = 1
		for j := 1; j < i; j++ {
			ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
		}
	}

	return ret
}
