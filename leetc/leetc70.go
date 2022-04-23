package leetc

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	left, right := 1, 1
	for i := 1; i < n; i++ {
		left, right = right, right+left
	}
	return right
}
