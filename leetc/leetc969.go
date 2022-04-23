package leetc

func PancakeSort(arr []int) []int {
	if len(arr) == 1 {
		return []int{}
	}
	idxMap := map[int]int{}
	for k, v := range arr {
		idxMap[v] = k
	}

	ret := make([]int, 0, len(arr)*2)
	for i := len(arr); i > 0; i-- {
		ret = pancakeReverse(idxMap[i], idxMap, arr, ret)
	}

	return ret
}

func pancakeReverse(i int, idxMap map[int]int, arr, ret []int) []int {
	q, j, k := arr[i]-1, i, 0
	ret = append(ret, i+1)
	for j > k {
		idxMap[arr[j]], idxMap[arr[k]] = idxMap[arr[k]], idxMap[arr[j]]
		arr[j], arr[k] = arr[k], arr[j]
		j--
		k++
	}

	k = 0
	ret = append(ret, arr[k])
	for q > k {
		idxMap[arr[q]], idxMap[arr[k]] = idxMap[arr[k]], idxMap[arr[q]]
		arr[q], arr[k] = arr[k], arr[q]
		q--
		k++
	}
	return ret
}
