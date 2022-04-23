package leetc

func longestPalindrome(s string) string {
	n, m := len(s), 1
	ret := s[0:1]
	left, right := 0, 0

	for i := 0; i < n; i++ {
		left = i - 1
		right = i + 1
		for left >= 0 && right < n {
			if s[left] != s[right] {
				break
			}
			if right-left+1 > m {
				ret = s[left : right+1]
				m = right - left + 1
			}
			left--
			right++
		}
		left = i
		right = i + 1
		for left >= 0 && right < n {
			if s[left] != s[right] {
				break
			}
			if right-left+1 > m {
				ret = s[left : right+1]
				m = right - left + 1
			}
			left--
			right++
		}
	}

	return ret
}
