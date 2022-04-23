package leetc

type MyLinkedList struct {
	left  *myNode
	right *myNode
	size  int
}

type myNode struct {
	Val  int
	next *myNode
	prev *myNode
}

func Constructor707() MyLinkedList {
	list := MyLinkedList{left: &myNode{}, right: &myNode{}}
	list.left.next = list.right
	//list.left.prev = &list.right
	list.right.prev = list.left
	//list.right.next = &list.left
	return list
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	var cur *myNode
	if index < this.size/2 {
		cur = this.left
		for i := index; i >= 0; i-- {
			cur = cur.next
		}
	} else {
		cur = this.right
		index = this.size - index
		for i := index; i > 0; i-- {
			cur = cur.prev
		}
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		add(this.left, val)
		this.size++
	} else if index > this.size {
		return
	} else {
		var cur *myNode
		if index <= this.size/2 {
			cur = this.left
			for i := index - 1; i >= 0; i-- {
				cur = cur.next
			}
			add(cur, val)
		} else {
			cur = this.right
			index = this.size - index
			for i := index - 1; i >= 0; i-- {
				cur = cur.prev
			}
			addPrev(cur, val)
		}
		this.size++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	var cur *myNode
	if index <= this.size/2 {
		cur = this.left
		for i := index - 1; i >= 0; i-- {
			cur = cur.next
		}
		deleteNode(cur)
	} else {
		cur = this.right
		index = this.size - index
		for i := index - 1; i >= 0; i-- {
			cur = cur.prev
		}
		deletePrevNode(cur)
	}
	this.size--
}

func add(this *myNode, val int) {
	next := this.next
	node := myNode{Val: val, next: next, prev: this}
	this.next = &node
	next.prev = &node
}

func addPrev(this *myNode, val int) {
	prev := this.prev
	node := myNode{Val: val, next: this, prev: prev}
	this.prev = &node
	prev.next = &node
}

func deleteNode(this *myNode) {
	this.next = this.next.next
	this.next.prev = this
}
func deletePrevNode(this *myNode) {
	this.prev = this.prev.prev
	this.prev.next = this
}
