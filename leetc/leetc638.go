package leetc

func ShoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(price)
	needP := map[string]int{}

	var dfs func([]byte) int
	dfs = func(curNeeds []byte) (minP int) {
		if v, has := needP[string(curNeeds)]; has {
			return v
		}
		for i, v := range curNeeds {
			minP += price[i] * int(v)
		}

		nextNeeds := make([]byte, n)
	outer:
		for _, spec := range special {
			for j, specNeed := range spec[:n] {
				if curNeeds[j] < byte(specNeed) {
					continue outer
				}
				nextNeeds[j] = curNeeds[j] - byte(specNeed)
			}
			minP = min638(minP, dfs(nextNeeds)+spec[n])
		}
		needP[string(curNeeds)] = minP
		return minP
	}

	curNeeds := make([]byte, n)
	for i, need := range needs {
		curNeeds[i] = byte(need)
	}

	return dfs(curNeeds)
}

func min638(a, b int) int {
	if a < b {
		return a
	}
	return b
}
