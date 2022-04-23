package leetc

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	res, slow, fast := head, head, head

	for {
		if fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			return nil
		}
		slow = slow.Next
		if slow == fast {
			for res != slow {
				res = res.Next
				slow = slow.Next
			}
			return res
		}
	}
}
