package leetc

func countEven(num int) int {
	var oddEven []int
	if num < 10 {
		oddEven = make([]int, 10)
	} else {
		oddEven = make([]int, num+1)
	}
	oddEven[1] = 1
	oddEven[3] = 1
	oddEven[5] = 1
	oddEven[7] = 1
	oddEven[9] = 1

	ret := 0
	for i := 1; i <= num; i++ {
		mod := i % 10
		j := i / 10
		if oddEven[mod]+oddEven[j] == 1 {
			oddEven[i] = 1
		} else {
			ret++
		}
	}

	return ret
}
