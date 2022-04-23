package leetc

type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor232() MyQueue {
	return MyQueue{make([]int, 0), make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		this.inToOut()
	}
	ret := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return ret
}

func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		this.inToOut()
	}
	return this.outStack[len(this.outStack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.inStack)+len(this.outStack) == 0
}

func (this *MyQueue) inToOut() {
	if len(this.inStack) == 0 {
		return
	}
	outS := make([]int, 0)
	for i := len(this.inStack) - 1; i >= 0; i-- {
		outS = append(outS, this.inStack[i])
	}
	this.inStack = make([]int, 0)
	this.outStack = outS
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
