package leetc

func maxPathSum(root *TreeNode) (res int) {
	res = root.Val
	var recur func(*TreeNode) int
	recur = func(root *TreeNode) int {
		if root == nil {
			return -1001
		}
		leftVal, rightVal := recur(root.Left), recur(root.Right)
		childMax := max124(leftVal, rightVal)
		if leftVal != -1001 && rightVal != -1001 {
			res = max124(res, leftVal+rightVal+root.Val)
		}

		rootVal := 0
		if childMax <= 0 {
			rootVal = root.Val
		} else {
			rootVal = root.Val + childMax
		}
		res = max124(res, rootVal)
		return rootVal
	}
	recur(root)

	return
}

func max124(a, b int) int {
	if a > b {
		return a
	}
	return b
}
