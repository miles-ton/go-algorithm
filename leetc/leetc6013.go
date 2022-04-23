package leetc

func mergeNodes(head *ListNode) *ListNode {
	cur := head.Next
	head = cur

	for cur != nil && cur.Next != nil {
		tmp := cur.Next
		if tmp.Val == 0 {
			cur.Next = tmp.Next
			cur = cur.Next
		} else {
			cur.Val += tmp.Val
			cur.Next = tmp.Next
		}
	}

	return head
}
