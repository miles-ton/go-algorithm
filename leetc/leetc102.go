package leetc

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	var nextLv, thisLv []*TreeNode
	thisLv = append(thisLv, root)
	for len(thisLv) != 0 {
		elem := make([]int, 0, len(thisLv))
		for _, node := range thisLv {
			elem = append(elem, node.Val)
			if node.Left != nil {
				nextLv = append(nextLv, node.Left)
			}
			if node.Right != nil {
				nextLv = append(nextLv, node.Right)
			}
		}
		thisLv = nextLv
		nextLv = nil
		res = append(res, elem)
	}

	return res
}
