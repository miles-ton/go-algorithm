package leetc

import "math"

type MinStack struct {
	stack []int
	min   int
}

func Constructor155() MinStack {
	return MinStack{make([]int, 0), math.MaxInt32}
}

func (this *MinStack) Push(val int) {
	if val <= this.min {
		this.stack = append(this.stack, this.min)
		this.min = val
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	top := this.stack[len(this.stack)-1]
	if top == this.min {
		this.stack = this.stack[:len(this.stack)-1]
		this.min = this.stack[len(this.stack)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
