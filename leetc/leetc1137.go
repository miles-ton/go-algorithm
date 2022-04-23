package leetc

func tribonacci(n int) int {
	var dp []int
	if n < 2 {
		dp = make([]int, 3)
	} else {
		dp = make([]int, n+1)
	}
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}
