package leetc

func minimumTime(s string) int {
	l := len(s)
	prev, ret := 0, l
	for i, v := range s {
		if v == '1' {
			prev = min2167(prev+2, i+1)
		}
		ret = min2167(ret, prev+l-1-i)
	}
	return ret
}

func min2167(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
