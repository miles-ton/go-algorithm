package leetc

func maxProfit309(prices []int) int {
	dp := make([][]int, 3)
	for i := range dp {
		dp[i] = make([]int, len(prices))
	}
	dp[0][0] = -prices[0]
	dp[1][0] = 0
	dp[2][0] = 0

	for i, v := range prices[1:] {
		dp[0][i+1] = max309(dp[0][i], dp[2][i]-v)
		dp[1][i+1] = dp[0][i] + v
		dp[2][i+1] = max309(dp[1][i], dp[2][i])
	}

	return max309(max309(dp[0][len(prices)-1], dp[1][len(prices)-1]), dp[2][len(prices)-1])
}

func max309(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
