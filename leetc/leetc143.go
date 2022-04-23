package leetc

var first *ListNode

func ReorderList(head *ListNode) {
	first = head
	recur143(head)
}

func recur143(second *ListNode) {
	if second.Next == nil {
		return
	}
	recur143(second.Next)
	if first == nil || first.Next == nil || first == second {
		first = nil
		return
	}

	tmp := first.Next
	first.Next = second.Next
	second.Next = nil
	first.Next.Next = tmp
	first = tmp
}
