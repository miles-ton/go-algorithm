package leetc

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	left, right := head, head.Next
	for right != nil && right.Next != nil {
		tmp := left.Next
		left.Next = right.Next
		right.Next = right.Next.Next
		right = right.Next

		left = left.Next
		left.Next = tmp
	}

	return head
}
