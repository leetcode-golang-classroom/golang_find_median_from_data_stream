# golang_find_median_from_data_stream

The **median** is the middle value in an ordered integer list. If the size of the list is even, there is no middle value and the median is the mean of the two middle values.

- For example, for `arr = [2,3,4]`, the median is `3`.
- For example, for `arr = [2,3]`, the median is `(2 + 3) / 2 = 2.5`.

Implement the MedianFinder class:

- `MedianFinder()` initializes the `MedianFinder` object.
- `void addNum(int num)` adds the integer `num` from the data stream to the data structure.
- `double findMedian()` returns the median of all elements so far. Answers within `10^{5}` of the actual answer will be accepted.

## Examples

**Example 1:**

```
Input
["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
[[], [1], [2], [], [3], []]
Output
[null, null, null, 1.5, null, 2.0]

Explanation
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1);    // arr = [1]
medianFinder.addNum(2);    // arr = [1, 2]
medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
medianFinder.addNum(3);    // arr[1, 2, 3]
medianFinder.findMedian(); // return 2.0
```

**Constraints:**

- $`10^{-5}$ <= num <= $10^5$`
- There will be at least one element in the data structure before calling `findMedian`.
- At most $5*10^4$ calls will be made to `addNum` and `findMedian`.

**Follow up:**

- If all integer numbers from the stream are in the range `[0, 100]`, how would you optimize your solution?
- If `99%` of all integer numbers from the stream are in the range `[0, 100]`, how would you optimize your solution?

## 解析

題目要設計一個資料結構 MedianFinder 需要具有以下 method

1. Constructor: 初始化 MedianFinder 結構
2. void addNum(num int): 用來新增 數字到目前陣列資料
3. double findMedium(): 用來找出陣列中的中位數

直覺作法是用一個陣列 arr 來儲存所有輸入的數值

每次新增數字到陣列 arr 時都照大小順序輸入

假設使用 BinarySearch 找出該數值該放入的位置arr 這樣就是 O(logN)

而中位數就是 就是$(arr[Math.floor((arr.length-1)/2] + arr[Math.floor(arr.length/2)])/2$

另一個作法就是使用兩個長度相差1 的 Heap來儲存

MinHeap 來放較大的一半

MaxHeap 來放較小的一半

當 MinHeap 長度大於 MaxHeap 代表 MinHeap 的 root 就是那個中位數

當 MaxHeap 長度大於 MinHeap 代表 MaxHeap 的 root 就是那個中位數

當長度相等時，中位數就是兩個 Heap的 root 相加除以 2

每次新增數值時，先把值新增到 MinHeap

然後檢查 MinHeap Root 是否小於 MaxHeap 的 Root

如果不是 則把 MinHeap Pop 出來的值 Push 到 MaxHeap

檢查 MinHeap 長度是否 > MaxHeap 長度 + 1

如果是 ⇒ 則把 MinHeap Pop 出來的值 Push 到 MaxHeap

檢查 MaxHeap 長度是否 > MinHeap 長度 + 1

如果是 ⇒ 則把 MaxHeap Pop 出來的值 Push 到 MinHeap

![](https://i.imgur.com/rUqrZ0a.png)

## 程式碼
```go
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

```
## 困難點

1. Understand MinHeap and MaxHeap
2. Understand timing and process of modify MinHeap to MaxHeap

## Solve Point

- [x]  Understand what problem need to solve
- [x]  Analysis Complexity