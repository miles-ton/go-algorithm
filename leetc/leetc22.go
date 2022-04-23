package leetc

func generateParenthesis(n int) []string {

	var ret []string
	var dfs func(liveLeft, availLeft, availRight, curI int, braces []byte)
	dfs = func(liveLeft, availLeft, availRight, curI int, braces []byte) {
		if availRight == 0 {
			ret = append(ret, string(braces))
			return
		}
		if availLeft > 0 {
			braces[curI] = '('
			dfs(liveLeft+1, availLeft-1, availRight, curI+1, braces)
		}
		if availRight > 0 && liveLeft > 0 {
			braces[curI] = ')'
			dfs(liveLeft-1, availLeft, availRight-1, curI+1, braces)
		}
	}
	dfs(0, n, n, 0, make([]byte, n*2))

	return ret
}
