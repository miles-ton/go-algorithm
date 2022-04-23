package leetc

func merge(nums1 []int, m int, nums2 []int, n int) {
	tmp1 := make([]int, m)
	copy(tmp1, nums1[:m])
	i1, i2 := 0, 0
	for i := range nums1 {
		if i1 < m && i2 < n {
			if tmp1[i1] < nums2[i2] {
				nums1[i] = tmp1[i1]
				i1++
			} else {
				nums1[i] = nums2[i2]
				i2++
			}
		} else {
			if i1 < m {
				copy(nums1[i:len(nums1)], tmp1[i1:])
			} else {
				copy(nums1[i:len(nums1)], nums2[i2:])
			}
			break
		}
	}
}
