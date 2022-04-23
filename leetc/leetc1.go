package leetc

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, v := range nums {
		hash[v] = i
	}

	for i, v := range nums {
		m, ok := hash[target-v]
		if ok && i != m {
			return []int{i, m}
		}
	}
	return []int{1, 2}
}
