package leetc

func SortList(head *ListNode) *ListNode {
	length, subLen := 0, 1
	h := head
	for h != nil {
		length++
		h = h.Next
	}
	ret := new(ListNode)
	ret.Next = head
	for subLen < length {
		prevNode := ret
		for prevNode.Next != nil {
			left, right := prevNode, prevNode
			for tmp := subLen; tmp > 0; tmp-- {
				if prevNode.Next != nil {
					prevNode = prevNode.Next
				} else {
					break
				}
			}
			right = prevNode
			for tmp := subLen; tmp > 0; tmp-- {
				if prevNode.Next != nil {
					prevNode = prevNode.Next
				} else {
					break
				}
			}
			subTail := prevNode.Next
			prevNode.Next = nil
			lb := right.Next
			right.Next = nil
			mergedHead, mergedTail := mergeLink(left.Next, lb)
			left.Next = mergedHead
			mergedTail.Next = subTail
			prevNode = mergedTail
		}
		subLen *= 2
	}

	return ret.Next
}

func mergeLink(la, lb *ListNode) (head, tail *ListNode) {
	empHead := new(ListNode)
	cur := empHead

	for la != nil && lb != nil {
		if la.Val <= lb.Val {
			cur.Next = la
			la = la.Next
		} else {
			cur.Next = lb
			lb = lb.Next
		}
		cur = cur.Next
	}

	if la == nil {
		cur.Next = lb
	} else {
		cur.Next = la
	}
	for cur.Next != nil {
		cur = cur.Next
	}

	return empHead.Next, cur
}
