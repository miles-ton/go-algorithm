package leetc

import "sort"

func sortEvenOdd(nums []int) []int {
	numsLen := len(nums)
	odd := make([]int, numsLen/2)
	var even []int
	if numsLen%2 == 1 {
		even = make([]int, numsLen/2+1)
	} else {
		even = make([]int, numsLen/2)
	}
	for i := range even {
		even[i] = nums[i*2]
	}
	for i := range odd {
		odd[i] = nums[i*2+1]
	}

	sort.Sort(sort.Reverse(mySortRule(odd)))
	sort.Ints(even)

	for i := range even {
		nums[i*2] = even[i]
	}
	for i := range odd {
		nums[i*2+1] = odd[i]
	}

	return nums
}

type mySortRule []int

func (my mySortRule) Len() int {
	return len(my)
}

func (my mySortRule) Swap(i, j int) {
	my[i], my[j] = my[j], my[i]
}

func (my mySortRule) Less(i, j int) bool {
	return my[i] < my[j]
}
