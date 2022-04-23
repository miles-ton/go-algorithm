package leetc

import "container/heap"

func nthUglyNumber(n int) int {
	ugly := &IntHeap{1}
	factors := []int{2, 3, 5}
	dup := map[int]struct{}{}
	cur := (heap.Pop(ugly)).(int)
	for i := 1; i != n; cur = (heap.Pop(ugly)).(int) {
		for _, f := range factors {
			m := cur * f
			if _, exist := dup[m]; exist {
				continue
			}
			dup[m] = struct{}{}
			heap.Push(ugly, m)
		}
		i++
	}
	return cur
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
