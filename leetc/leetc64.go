package leetc

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	for i, v := range grid[0][1:] {
		dp[0][i+1] = dp[0][i] + v
	}
	for i := range grid[1:] {
		dp[i+1][0] = dp[i][0] + grid[i+1][0]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = min64(dp[i][j-1], dp[i-1][j]) + grid[i][j]
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}

func min64(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
