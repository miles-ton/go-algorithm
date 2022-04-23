package leetc

func lengthOfLongestSubstring(s string) int {
	subStr := make(map[byte]int)

	ret := 0
	j, i := 0, 0
	for ; i < len(s); i++ {
		cur := s[i]
		if mapContains(subStr, cur) {
			newJ := subStr[cur]
			for ; j < newJ; j++ {
				delete(subStr, s[j])
			}
			subStr[cur] = i
			j = newJ + 1
			continue
		}
		subStr[cur] = i
		ret = max3(len(subStr), ret)
	}
	return ret
}

func mapContains(m map[byte]int, t byte) bool {
	_, ok := m[t]
	return ok
}

func max3(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
