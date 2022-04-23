package leetc

var ps = []byte{'a', 'b', 'c'}

func modifyString(s string) string {
	cs := []byte(s)
	if len(cs) == 1 && cs[0] == '?' {
		return "a"
	}
	if cs[0] == '?' {
		cs[0] = pick('?', cs[1])
	}
	if cs[len(s)-1] == '?' {
		cs[len(s)-1] = pick(cs[len(s)-2], '?')
	}
	for i := 1; i < len(cs)-1; i++ {
		if cs[i] == '?' {
			cs[i] = pick(cs[i-1], cs[i+1])
		}
	}
	return string(cs)
}

func pick(a, b byte) byte {
	for _, c := range ps {
		if c != a && c != b {
			return c
		}
	}
	return 'a'
}
