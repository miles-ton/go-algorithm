package leetc

func minSteps(n int) (res int) {
	res = 1
	n--

	var recur func(int, int) bool

	recur = func(al int, ac int) bool {
		if (al*ac)%(al*2) != 0 {
			return false
		}
		al *= 2
		ac /= 2
		res += 2
		ac--

		for ac > 0 && !recur(al, ac) {
			ac--
			res++
		}
		return true
	}

	recur(1, n-1)

	return res
}
