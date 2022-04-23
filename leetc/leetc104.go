package leetc

func maxDepth(root *TreeNode) (res int) {
	if root == nil {
		return 0
	}
	return recur104(root, 1)
}

func recur104(root *TreeNode, depth int) int {
	leftDepth, rightDepth := depth, depth
	if root.Left != nil {
		leftDepth = recur104(root.Left, depth+1)
	}
	if root.Right != nil {
		rightDepth = recur104(root.Right, depth+1)
	}
	return max104(leftDepth, rightDepth)
}

func max104(a, b int) int {
	if a > b {
		return a
	}
	return b
}
