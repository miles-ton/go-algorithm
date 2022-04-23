package leetc

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	ret := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min221(min221(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
				ret = max221(ret, dp[i][j])
			} else {
				dp[i][j] = 0
			}
		}
	}
	return ret * ret
}

func max221(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min221(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
