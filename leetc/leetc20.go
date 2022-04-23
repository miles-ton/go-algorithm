package leetc

func isValid(s string) bool {
	var stack []rune
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			n := len(stack)
			if c == ')' {
				if len(stack) > 0 && stack[n-1] == '(' {
					stack = stack[:n-1]
				} else {
					return false
				}
			} else if c == ']' {
				if len(stack) > 0 && stack[n-1] == '[' {
					stack = stack[:n-1]
				} else {
					return false
				}
			} else {
				if len(stack) > 0 && stack[n-1] == '{' {
					stack = stack[:n-1]
				} else {
					return false
				}
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
