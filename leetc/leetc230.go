package leetc

func kthSmallest(root *TreeNode, k int) (ret int) {

	var mr func(*TreeNode)
	ret = -1
	mr = func(root *TreeNode) {
		if root == nil || ret != -1 {
			return
		}
		mr(root.Left)
		k--
		if k == 0 {
			ret = root.Val
		}
		mr(root.Right)
	}
	mr(root)

	return ret
}
