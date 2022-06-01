package sol

import "container/heap"

type MaxHeap []int
type MinHeap []int

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	small MaxHeap
	large MinHeap
}

func Constructor() MedianFinder {
	maxHeap := MaxHeap{}
	minHeap := MinHeap{}
	heap.Init(&maxHeap)
	heap.Init(&minHeap)
	return MedianFinder{
		small: maxHeap,
		large: minHeap,
	}
}

func (this *MedianFinder) AddNum(num int) {
	// push val into small first
	heap.Push(&this.small, num)
	// check if small[0] < large[0]
	if this.small.Len() > 0 && this.large.Len() > 0 && this.small[0] > this.large[0] {
		// pop small, push into large
		max := heap.Pop(&this.small).(int)
		heap.Push(&this.large, max)
	}
	// check if this.small.Len() > this.large.Len() + 1
	if this.small.Len() > this.large.Len()+1 {
		max := heap.Pop(&this.small).(int)
		heap.Push(&this.large, max)
	}
	// check if this.large.Len() > this.small.Len() + 1
	if this.large.Len() > this.small.Len()+1 {
		min := heap.Pop(&this.large).(int)
		heap.Push(&this.small, min)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.small.Len() > this.large.Len() {
		return float64(this.small[0])
	}
	if this.small.Len() < this.large.Len() {
		return float64(this.large[0])
	}
	return float64(this.small[0]+this.large[0]) / 2
}

/**
* Your MedianFinder object will be instantiated and called as such:
* obj := Constructor();
* obj.AddNum(num);
* param_2 := obj.FindMedian();
 */
