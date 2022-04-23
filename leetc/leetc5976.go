package leetc

func checkValid(matrix [][]int) bool {
	dup := make([]bool, len(matrix))
	for _, row := range matrix {
		for _, v := range row {
			if dup[v-1] {
				return false
			}
			dup[v-1] = true
		}
		cleanSlice(dup)
	}
	l := len(matrix)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if dup[matrix[j][i]-1] {
				return false
			}
			dup[matrix[j][i]-1] = true
		}
		cleanSlice(dup)
	}
	return true
}

func cleanSlice(s []bool) {
	for i := range s {
		s[i] = false
	}
}
