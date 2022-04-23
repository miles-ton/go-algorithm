package leetc

import "sort"

func smallestNumber(num int64) int64 {
	var flag int64
	if num > 0 {
		flag = 1
	} else if num < 0 {
		flag = -1
	} else {
		return 0
	}
	numArr := make([]int, 0, 40)
	for num != 0 {
		numArr = append(numArr, int(num%10))
		num /= 10
	}
	sort.Ints(numArr)
	var ret int64
	fNoZeroI := 0
	for ; fNoZeroI < len(numArr); fNoZeroI++ {
		if numArr[fNoZeroI] == 0 {
			continue
		} else {
			break
		}
	}
	numArr[0], numArr[fNoZeroI] = numArr[fNoZeroI], numArr[0]
	for _, v := range numArr {
		ret *= 10
		ret += int64(v)
	}
	if flag == -1 && ret >= 0 {
		ret *= -1
	}

	return ret
}
