package leetc

func maxProfit714(prices []int, fee int) int {
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, len(prices))
	}
	dp[0][0] = -prices[0]
	for i, v := range prices[1:] {
		dp[0][i+1] = max714(dp[0][i], dp[1][i]-v)
		dp[1][i+1] = max714(dp[0][i]+v-fee, dp[1][i])
	}

	return dp[1][len(prices)-1]
}

func max714(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
