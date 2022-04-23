package leetc

func postorderTraversal(root *TreeNode) []int {

	return postOrder(root, []int{})
}

func postOrder(root *TreeNode, ret []int) []int {
	if root == nil {
		return ret
	}

	ret = postOrder(root.Left, ret)
	ret = postOrder(root.Right, ret)

	return append(ret, root.Val)
}
