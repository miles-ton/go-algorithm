package leetc

func longestCommonSubsequence1143(text1 string, text2 string) int {
	bs1 := []byte(text1)
	bs2 := []byte(text2)
	dp := make([][]int, len(bs1)+1)
	dp[0] = make([]int, len(bs2)+1)
	for i := range bs1 {
		dp[i+1] = make([]int, len(bs2)+1)
	}

	for i, v := range bs1 {
		for i2, v2 := range bs2 {
			if v2 == v {
				dp[i+1][i2+1] = dp[i][i2] + 1
			} else {
				dp[i+1][i2+1] = max1143(dp[i][i2+1], dp[i+1][i2])
			}
		}
	}

	return dp[len(bs1)][len(bs2)]
}

func max1143(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
