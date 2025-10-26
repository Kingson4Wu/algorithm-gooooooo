package _02510

/*
*
根据回忆独立完成

执行用时分布
25
ms
击败
73.68%
复杂度分析
消耗内存分布
16.73
MB
击败
21.05%
*/
type heap struct {
	arr     []int
	min     bool // 是否最小堆
	compare func(i, j int) bool
}

func NewHeap(min bool) *heap {
	h := &heap{}
	if min {
		h.compare = func(i, j int) bool {
			return h.arr[i] < h.arr[j]
		}
	} else {
		h.compare = func(i, j int) bool {
			return h.arr[i] > h.arr[j]
		}
	}
	return h
}

// 从下往上调整
func (h *heap) adjustUp(index int) {
	root := (index+1)/2 - 1
	for root >= 0 {
		if h.compare(root, index) {
			break
		}
		h.arr[root], h.arr[index] = h.arr[index], h.arr[root]
		index = root
		root = (index+1)/2 - 1
	}
}
func (h *heap) adjustDown(root int) {
	index := root*2 + 1
	for index < len(h.arr) {
		if index+1 < len(h.arr) && h.compare(index+1, index) {
			index++
		}
		if h.compare(root, index) {
			break
		}
		h.arr[root], h.arr[index] = h.arr[index], h.arr[root]
		root = index
		index = root*2 + 1
	}
}

func (h *heap) Push(num int) {
	h.arr = append(h.arr, num)
	h.adjustUp(len(h.arr) - 1)
}

func (h *heap) Pop() int {
	if len(h.arr) == 0 {
		return 0
	}
	v := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.adjustDown(0)
	return v
}

func (h *heap) Top() int {
	if len(h.arr) == 0 {
		return 0
	}
	return h.arr[0]
}

func (h *heap) Len() int {
	return len(h.arr)
}

type MedianFinder struct {
	minHeap *heap
	maxHeap *heap
}

/** initialize your data structure here. */
func Constructor1() MedianFinder {
	return MedianFinder{
		minHeap: NewHeap(true),
		maxHeap: NewHeap(false),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.minHeap.Top() <= num {
		this.minHeap.Push(num)
	} else {
		this.maxHeap.Push(num)
	}
	if this.minHeap.Len()-this.maxHeap.Len() > 1 {
		this.maxHeap.Push(this.minHeap.Pop())
	} else if this.maxHeap.Len()-this.minHeap.Len() > 1 {
		this.minHeap.Push(this.maxHeap.Pop())
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Len() > this.maxHeap.Len() {
		return float64(this.minHeap.Top())
	} else if this.maxHeap.Len() > this.minHeap.Len() {
		return float64(this.maxHeap.Top())
	} else {
		return float64(this.maxHeap.Top()+this.minHeap.Top()) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
