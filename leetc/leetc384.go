package leetc

import "math/rand"

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head}
}

func (receiver *Solution) GetRandom() int {
	ret := receiver.head.Val
	start := receiver.head.Next
	count := 2

	for start != nil {
		if rand.Intn(count) == 0 {
			ret = start.Val
		}
		start = start.Next
		count++
	}
	return ret
}
