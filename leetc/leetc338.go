package leetc

func countBits(n int) []int {
	if n == 0 {
		return []int{}
	}
	cur, ret := 1, make([]int, n+1)
	ret[0] = 0
	for i := 1; cur <= n; i <<= 1 {
		for j := 0; j < i && cur <= n; j++ {
			cur++
			ret[i+j] = ret[j] + 1
		}
	}
	return ret
}
