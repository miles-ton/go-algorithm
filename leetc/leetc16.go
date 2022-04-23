package leetc

import "sort"

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	ret := nums[0] + nums[1] + nums[n-1]
	for i := 0; i < n; i++ {
		for j, k := i+1, n-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			} else if sum < target {
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}
			} else {
				for k--; j < k && nums[k] == nums[k+1]; k-- {
				}
			}
			ret = min16(ret, sum, target)
		}
	}
	return ret
}

func min16(a, b, tar int) int {
	if abs16(tar-a) < abs16(tar-b) {
		return a
	}
	return b
}

func abs16(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
