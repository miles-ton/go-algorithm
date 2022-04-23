package leetc

import "sort"

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return [][]int{}
	}
	sort.Ints(nums)

	var ret [][]int
	for i := 0; i < n-3; {
		for j := i + 1; j < n-2; {
			for k, q := j+1, n-1; k < n-1 && k < q; {
				sum := nums[i] + nums[j] + nums[k] + nums[q]
				if sum == target {
					ret = append(ret, []int{nums[i], nums[j], nums[k], nums[q]})
					for k++; k < q && nums[k] == nums[k-1]; k++ {
					}
					for q--; q > k && nums[q] == nums[q+1]; q-- {
					}
				} else if sum < target {
					for k++; k < q && nums[k] == nums[k-1]; k++ {
					}
				} else {
					for q--; q > k && nums[q] == nums[q+1]; q-- {
					}
				}
			}
			for j++; j < n-2 && nums[j] == nums[j-1]; j++ {
			}
		}
		for i++; i < n-3 && nums[i] == nums[i-1]; i++ {
		}
	}
	return ret
}
