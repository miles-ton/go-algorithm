package leetc

import (
	"strconv"
	"strings"
)

func optimalDivision(nums []int) string {
	n := len(nums)
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	if len(nums) == 2 {
		return strconv.Itoa(nums[0]) + "/" + strconv.Itoa(nums[1])
	}
	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(nums[0]))
	sb.WriteString("/(")
	for _, v := range nums[1 : len(nums)-1] {
		sb.WriteString(strconv.Itoa(v))
		sb.WriteRune('/')
	}
	sb.WriteString(strconv.Itoa(nums[n-1]))
	sb.WriteString(")")
	return sb.String()
}
