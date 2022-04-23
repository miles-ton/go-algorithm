package leetc

import "math"

func minFallingPathSum(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+2)
		dp[i][0] = math.MaxInt32
		dp[i][n+1] = math.MaxInt32
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = min931(min931(dp[i-1][j-1], dp[i-1][j]), dp[i-1][j+1]) + matrix[i-1][j-1]
		}
	}
	ret := math.MaxInt32
	for _,v := range dp[m][1 : n+1] {
		ret = min931(ret, v)
	}
	return ret
}

func min931(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
