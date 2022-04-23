package leetc

func KnightProbability(n int, k int, row int, column int) float64 {
	if k == 0 {
		return 1.0
	}
	curRates := make([][]float64, n)
	lastRates := make([][]float64, n)
	for i := range curRates {
		curRates[i] = make([]float64, n)
		lastRates[i] = make([]float64, n)
	}
	routes := [][]int{{-2, -1}, {-1, -2}, {1, -2}, {2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}}
	for j := 0; j < n; j++ {
		for q := 0; q < n; q++ {
			for _, route := range routes {
				if j+route[0] >= 0 && j+route[0] < n && q+route[1] >= 0 && q+route[1] < n {
					curRates[j][q] += float64(1) / 8
				}
			}
		}
	}
	for i := 0; i < k; i++ {
		lastRates = curRates
		curRates = make([][]float64, n)
		for i := range curRates {
			curRates[i] = make([]float64, n)
		}
		for j := 0; j < n; j++ {
			for q := 0; q < n; q++ {
				for _, route := range routes {
					if j+route[0] >= 0 && j+route[0] < n && q+route[1] >= 0 && q+route[1] < n {
						curRates[j][q] += lastRates[j+route[0]][q+route[1]] / 8
					}
				}
			}
		}
	}

	return lastRates[row][column]
}
