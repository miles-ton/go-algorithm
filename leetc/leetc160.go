package leetc

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ha, hb := headA, headB
	for ha != nil || hb != nil {
		if ha == nil {
			ha = headB
		}
		if hb == nil {
			hb = headA
		}
		if ha == hb {
			return ha
		}
		ha = ha.Next
		hb = hb.Next
	}
	return nil
}
