package leetc

import "math"

func myAtoi(s string) int {
	str := []byte(s[:])
	ret, cur, lennn := 0, 0, len(str)

	for _, v := range str {
		if v != ' ' {
			break
		}
		cur++
	}

	symbol := 1
	if cur < lennn && str[cur] == '-' {
		symbol = -1
		cur++
	} else if cur < lennn && str[cur] == '+' {
		cur++
	}

	for cur < lennn && str[cur] == '0' {
		cur++
	}

	for ; cur < lennn && isDigit(str[cur]); cur++ {
		ret *= 10
		ret += symbol * btoi(str[cur])
		if ret < math.MinInt32 {
			return math.MinInt32
		} else if ret > math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return ret * symbol
}

func isDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

func btoi(b byte) int {
	return int(b - '0')
}
