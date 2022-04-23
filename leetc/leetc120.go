package leetc

import "math"

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, i+1)
		if i == 0 {
			dp[i][0] = triangle[0][0]
		} else {
			dp[i][0] = dp[i-1][0] + triangle[i][0]
			dp[i][i] = dp[i-1][i-1] + triangle[i][i]
		}
	}
	for i := 2; i < m; i++ {
		for j := 1; j < len(triangle[i])-1; j++ {
			dp[i][j] = min120(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
	}

	ret := math.MaxInt32
	for _, v := range dp[m-1] {
		ret = min120(ret, v)
	}
	return ret
}

func min120(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
