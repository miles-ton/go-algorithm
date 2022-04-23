package leetc

func insertionSortList(head *ListNode) *ListNode {
	ret := new(ListNode)

	for head != nil {
		cur := ret
		hCur := head
		head = head.Next
		hCur.Next = nil
		for cur != nil {
			if cur.Next == nil {
				cur.Next = hCur
				break
			} else if cur.Next.Val >= hCur.Val {
				hCur.Next = cur.Next
				cur.Next = hCur
				break
			} else {
				cur = cur.Next
			}
		}
	}
	return ret.Next
}
