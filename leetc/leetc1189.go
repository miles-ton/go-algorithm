package leetc

func maxNumberOfBalloons(text string) int {
	cnt := map[int32]int{}
	for _, v := range text {
		cnt[v]++
	}
	balloon := "balon"
	cnt['l'] /= 2
	cnt['o'] /= 2
	ret := 10001
	for _, v := range balloon {
		ret = min1189(cnt[v], ret)
	}
	return ret
}

func min1189(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
