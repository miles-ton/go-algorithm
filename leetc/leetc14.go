package leetc

func longestCommonPrefix(strs []string) string {
	n := len(strs[0])
	ret := make([]uint8, 0, n)

end:
	for i := 0; i < n; i++ {
		c := strs[0][i]
		for _, s := range strs[1:] {
			if i >= len(s) || c != s[i] {
				break end
			}
		}
		ret = append(ret, c)
	}
	return string(ret)
}
