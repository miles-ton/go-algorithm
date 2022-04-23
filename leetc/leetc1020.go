package leetc

var m, n int

func NumEnclaves(grid [][]int) int {
	m, n = len(grid), len(grid[0])

	for i := 0; i < n; i++ {
		visitDepth(0, i, grid)
		visitDepth(m-1, i, grid)
	}
	for i := 0; i < m; i++ {
		visitDepth(i, 0, grid)
		visitDepth(i, n-1, grid)
	}
	oneCnt := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				oneCnt++
			}
		}
	}

	return oneCnt
}

func visitDepth(i, j int, grid [][]int) {
	if grid[i][j] == 1 {
		grid[i][j] = 0
	} else {
		return
	}
	if isInGridAndNotVisited(i-1, j, grid) {
		visitDepth(i-1, j, grid)
	}
	if isInGridAndNotVisited(i+1, j, grid) {
		visitDepth(i+1, j, grid)
	}
	if isInGridAndNotVisited(i, j-1, grid) {
		visitDepth(i, j-1, grid)
	}
	if isInGridAndNotVisited(i, j+1, grid) {
		visitDepth(i, j+1, grid)
	}
}

func isInGridAndNotVisited(i, j int, grid [][]int) bool {
	if i > 0 && i < m && j > 0 && j < n && grid[i][j] == 1 {
		return true
	}
	return false
}
