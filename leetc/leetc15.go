package leetc

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	var ret [][]int
	if n < 3 {
		return ret
	}
	sort.Ints(nums)
	first, second, third := 0, 1, n-1

	for first = 0; first < n-2; {
		second = first + 1
		third = n - 1
		for second < third {
			sum := nums[first] + nums[second] + nums[third]
			if sum == 0 {
				ret = append(ret, []int{nums[first], nums[second], nums[third]})
				newSec, newThi := second+1, third-1
				for newSec < third && nums[newSec] == nums[second] {
					newSec++
				}
				for newThi > second && nums[newThi] == nums[third] {
					newThi--
				}
				second, third = newSec, newThi
			} else if sum < 0 {
				second++
			} else {
				third--
			}
		}
		newFir := first + 1
		for newFir < n-2 && nums[newFir] == nums[first] {
			newFir++
		}
		first = newFir
	}
	return ret
}
