package leetc

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	virtualHead := &ListNode{Next: head}
	head = virtualHead
	left, right := head, head
	for i := 0; i < n; i++ {
		right = right.Next
	}
	for right.Next != nil {
		right = right.Next
		left = left.Next
	}
	left.Next = left.Next.Next

	return head.Next
}
