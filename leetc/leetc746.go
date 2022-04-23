package leetc

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)

	for i := 2; i < n+1; i++ {
		dp[i] = min746(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

func min746(a, b int) int {
	if a < b {
		return a
	}
	return b
}
