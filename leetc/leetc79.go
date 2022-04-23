package leetc

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	wordB := []byte(word)
	trace := make([][]bool, m)
	for i := 0; i < m; i++ {
		trace[i] = make([]bool, n)
	}
	res := false
	var recur func(int, int, []byte)
	recur = func(curI, curJ int, wordB []byte) {
		if len(wordB) == 0 {
			res = true
			return
		}
		if curI >= m || curI < 0 || curJ >= n || curJ < 0 || trace[curI][curJ] == true || res {
			return
		}
		trace[curI][curJ] = true
		if board[curI][curJ] == wordB[0] {
			recur(curI+1, curJ, wordB[1:])
			recur(curI-1, curJ, wordB[1:])
			recur(curI, curJ+1, wordB[1:])
			recur(curI, curJ-1, wordB[1:])
		}
		trace[curI][curJ] = false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			recur(i, j, wordB)
		}
	}
	return res
}
