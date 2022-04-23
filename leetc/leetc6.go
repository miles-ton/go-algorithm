package leetc

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n := len(s)
	var ret []byte
	interval := numRows*2 - 2
	pointCnt := 0
	for i := 0; i < n; i += interval {
		pointCnt++
		ret = append(ret, s[i])
	}
	foffset, soffset := 1, interval-1
	for i := 1; i < numRows; i++ {
		for j := 0; j < pointCnt; j++ {
			startPoint := j * interval
			if startPoint+foffset < n {
				if foffset == soffset {
					ret = append(ret, s[startPoint+foffset])
				} else {
					ret = append(ret, s[startPoint+foffset])
					if startPoint+soffset < n {
						ret = append(ret, s[startPoint+soffset])
					}
				}
			}
		}
		foffset++
		soffset--
	}
	return string(ret)
}
