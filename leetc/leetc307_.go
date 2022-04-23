package leetc

type NumArray struct {
	nums []int
}

func Constructor307(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (this *NumArray) Update(index int, val int) {
	this.nums[index] = val

}

func (this *NumArray) SumRange(left int, right int) (res int) {
	for left <= right {
		res += this.nums[left]
		left++
	}
	return
}
