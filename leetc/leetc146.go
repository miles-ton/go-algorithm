package leetc

type LRUCache struct {
	capacity int
	cache    map[int]int
	nodeIdx  map[int]*prioNode
	head     *prioNode
	tail     *prioNode
}

type prioNode struct {
	key  int
	next *prioNode
}

func Constructor146(capacity int) LRUCache {
	n := LRUCache{capacity, make(map[int]int), make(map[int]*prioNode), nil, nil}
	n.head = new(prioNode)
	n.tail = n.head
	return n
}

func (this *LRUCache) Get(key int) int {
	ret, ok := this.cache[key]
	if !ok {
		return -1
	}
	preNode := this.nodeIdx[key]
	this.upNode(preNode)
	return ret
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.cache[key]
	if ok {
		this.cache[key] = value
		this.upNode(this.nodeIdx[key])
		return
	}
	if len(this.cache) == this.capacity {
		this.evict()
	}
	this.cache[key] = value
	this.nodeIdx[key] = this.tail
	this.tail.next = &prioNode{key, nil}
	this.tail = this.tail.next
}

func (this *LRUCache) upNode(preNode *prioNode) {
	this.tail.next = preNode.next
	this.nodeIdx[preNode.next.key] = this.tail
	this.tail = this.tail.next
	preNode.next = this.tail.next
	this.nodeIdx[this.tail.next.key] = preNode
	this.tail.next = nil
}

func (this *LRUCache) evict() {
	if this.head.next == nil {
		return
	}
	if this.head.next.next != nil {
		nextNodeVal := this.head.next.next.key
		this.nodeIdx[nextNodeVal] = this.head
	}
	if this.head.next == this.tail {
		this.tail = this.head
	}
	delete(this.cache, this.head.next.key)
	delete(this.nodeIdx, this.head.next.key)
	this.head.next = this.head.next.next
}
