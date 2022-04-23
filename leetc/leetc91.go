package leetc

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	left, right := 1, 1

	for i := 1; i < len(s); i++ {
		if s[i-1] == '0' && s[i] == '0' {
			return 0
		}
		one, two := 0, 0
		if s[i-1] == '1' {
			two = left
		} else if s[i-1] == '2' && s[i] >= '0' && s[i] <= '6' {
			two = left
		}

		if s[i] != '0' {
			one = right
		}

		left, right = right, one+two
	}

	return right
}
