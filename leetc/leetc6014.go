package leetc

func repeatLimitedString(s string, repeatLimit int) string {
	cnt := make([]int, 26)
	for _, v := range s {
		cnt[v-'a']++
	}
	ret := make([]byte, 0)
out:
	for i := 25; i >= 0; i-- {
		if cnt[i] <= 0 {
			continue
		}
		for cnt[i] != 0 {
			if cnt[i] <= repeatLimit {
				for k := cnt[i]; k > 0; k-- {
					ret = append(ret, byte('a'+i))
				}
				cnt[i] = 0
			} else {
				ci := nextComplement(i-1, cnt)
				for k := repeatLimit; k > 0; k-- {
					ret = append(ret, byte('a'+i))
				}
				cnt[i] -= repeatLimit
				if ci == -1 {
					continue out
				}
				ret = append(ret, byte('a'+ci))
				cnt[ci]--
			}
		}
	}

	return string(ret)
}

func nextComplement(end int, cnt []int) int {
	for i := end; i >= 0; i-- {
		if cnt[i] > 0 {
			return i
		}
	}
	return -1
}
