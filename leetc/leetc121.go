package leetc

func maxProfit(prices []int) int {
	profile := 0
	minTillNow := prices[0]
	for i := 1; i < len(prices); i++ {
		profile = max121(prices[i]-minTillNow, profile)
		minTillNow = min121(minTillNow, prices[i])
	}
	return profile
}

func max121(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min121(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
