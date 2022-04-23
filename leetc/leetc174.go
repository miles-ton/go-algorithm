package leetc

func CalculateMinimumHP(dungeon [][]int) int {
	dp := make([][]int, len(dungeon))
	for i := range dp {
		dp[i] = make([]int, len(dungeon[0]))
	}
	m, n := len(dungeon), len(dungeon[0])
	dp[m-1][n-1] = max174(1-dungeon[m-1][n-1], 1)
	for i := n - 2; i >= 0; i-- {
		dp[m-1][i] = max174(dp[m-1][i+1]-dungeon[m-1][i], 1)
	}
	for i := m - 2; i >= 0; i-- {
		dp[i][n-1] = max174(dp[i+1][n-1]-dungeon[i][n-1], 1)
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = max174(min174(dp[i+1][j], dp[i][j+1])-dungeon[i][j], 1)
		}
	}
	return dp[0][0]
}

func min174(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max174(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
