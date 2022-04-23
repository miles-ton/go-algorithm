package leetc

func lenLongestFibSubseq(arr []int) int {
	viMap := make(map[int]int)
	for i, v := range arr {
		viMap[v] = i
	}

	al := len(arr)
	dp := make([][]int, al)
	for i := range dp {
		dp[i] = make([]int, al)
	}
	ret := 0
	for i := range arr {
		for j := i - 1; j > 0; j-- {
			idx, ok := viMap[arr[i]-arr[j]]
			if ok && idx < j {
				dp[j][i] = dp[idx][j] + 1
				ret = max093(dp[j][i], ret)
			}
		}
	}
	if ret > 0 {
		ret += 2
	}
	return ret
}

func max093(a, b int) int {
	if a > b {
		return a
	}
	return b
}
