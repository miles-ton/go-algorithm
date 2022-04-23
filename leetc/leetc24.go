package leetc

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	next := head.Next
	head.Next = nil
	return recur(head, next)
}
func recur(last, cur *ListNode) *ListNode {
	if cur == nil {
		return last
	}
	next := cur.Next
	cur.Next = last
	return recur(cur, next)
}
