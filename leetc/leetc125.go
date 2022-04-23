package leetc

func IsPalindrome(s string) bool {
	head, tail := 0, len(s)-1

	for head < tail {
		for head < tail && !isLetter125(s[head]) && !isDigit125(s[head]) {
			head++
		}
		for head < tail && !isLetter125(s[tail]) && !isDigit125(s[tail]) {
			tail--
		}
		if head < tail && !equal(s[head], s[tail]) {
			return false
		}
		head++
		tail--
	}
	return true
}

func equal(a, b uint8) bool {
	if isLetter125(a) {
		if isLetter125(b) {
			if toLower125(a) == toLower125(b) {
				return true
			}
		}
	} else {
		if isDigit125(b) {
			return a == b
		}
	}
	return false
}

func isLetter125(l uint8) bool {
	return (l >= 'a' && l <= 'z') || (l >= 'A' && l <= 'Z')
}
func toLower125(l uint8) uint8 {
	if l >= 'A' && l <= 'Z' {
		return 'a' + l - 'A'
	}
	return l
}
func isDigit125(d uint8) bool {
	return d >= '0' && d <= '9'
}
