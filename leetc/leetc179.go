package leetc

import (
	"sort"
	"strconv"
)

type MySortRule []string

func (my MySortRule) Len() int {
	return len(my)
}

func (my MySortRule) Swap(i, j int) {
	my[i], my[j] = my[j], my[i]
}

func (my MySortRule) Less(i, j int) bool {
	var l int
	if len(my[i]) > len(my[j]) {
		l = len(my[j])
	} else {
		l = len(my[i])
	}
	for k := 0; k < l; k++ {
		if my[i][k] < my[j][k] {
			return true
		} else if my[i][k] > my[j][k] {
			return false
		}
	}
	return my[i]+my[j] < my[j]+my[j]
}

func largestNumber(nums []int) string {
	strNums := make(MySortRule, len(nums))
	for i := range strNums {
		strNums[i] = strconv.Itoa(nums[i])
	}
	sort.Sort(strNums)
	var ret string
	for i := len(strNums) - 1; i >= 0; i-- {
		ret += strNums[i]
	}
	if ret[0] == '0' {
		return "0"
	}
	return ret
}
