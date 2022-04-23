package leetc

func minDistance(word1 string, word2 string) int {
	w1, w2 := []byte(word1), []byte(word2)

	dp := make([][]int, len(w1))
	for i := range dp {
		dp[i] = make([]int, len(w2))
	}
	a := w1[0]
	for i := range w2 {
		if a == w2[i] {
			dp[0][i] = 1
		}
	}
	a = w2[0]
	for i := range w1 {
		if a == w1[i] {
			dp[i][0] = 1
		}
	}

	common := 0
	for i := 1; i < len(w1); i++ {
		for j := 1; j < len(w2); j++ {
			if w1[i] == w2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > common {
					common = dp[i][j]
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	return len(w1) - common + len(w2) - common
}
