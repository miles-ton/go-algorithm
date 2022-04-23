package leetc

func preorderTraversal(root *TreeNode) []int {
	return frontSeq(root, []int{})
}

func frontSeq(root *TreeNode, ret []int) []int {
	if root == nil {
		return ret
	}
	ret = append(ret, root.Val)
	ret = frontSeq(root.Left, ret)
	return frontSeq(root.Right, ret)
}
