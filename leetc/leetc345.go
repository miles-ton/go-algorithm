package leetc

func ReverseVowels(s string) string {
	vowels := map[uint8]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}
	ret := make([]uint8, len(s))
	head, tail := 0, len(s)-1
	for head <= tail {
		for ; head <= tail; head++ {
			_, ok := vowels[s[head]]
			if ok {
				break
			}
			ret[head] = s[head]
		}
		for ; head <= tail; tail-- {
			_, ok := vowels[s[tail]]
			if ok {
				break
			}
			ret[tail] = s[tail]
		}
		if head <= tail {
			ret[head] = s[tail]
			ret[tail] = s[head]
			head++
			tail--
		}
	}
	return string(ret)
}
