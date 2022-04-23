package leetc

func twoSum167(numbers []int, target int) []int {
	index1, index2 := 0, len(numbers)-1
	sum := numbers[index1] + numbers[index2]
	for sum != target {
		if sum > target {
			index2--
			sum = numbers[index1] + numbers[index2]
		} else if sum < target {
			index1++
			sum = numbers[index1] + numbers[index2]
		}
	}
	return []int{index1+1, index2+1}
}
