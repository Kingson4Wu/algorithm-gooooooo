package _02510

import (
	"fmt"
	"sort"
	"testing"
)

type heapNode struct {
	id      int
	endTime int
	count   int
}

type minHeap struct {
	arr     []*heapNode
	compare func(root, index *heapNode) bool
}

func NewMinHeap(arr []*heapNode, compare func(root, index *heapNode) bool) *minHeap {
	h := &minHeap{
		arr:     arr,
		compare: compare,
	}
	for i := len(h.arr)/2 - 1; i >= 0; i-- {
		h.adjustTopToBottom(i)
	}
	return h
}

func (h *minHeap) Push(node *heapNode) {
	h.arr = append(h.arr, node)
	h.adjustBottomToTop(len(h.arr) - 1)
}

func (h *minHeap) Pop() *heapNode {
	node := h.arr[0]
	h.arr[0], h.arr[len(h.arr)-1] = h.arr[len(h.arr)-1], h.arr[0]
	h.arr = h.arr[:len(h.arr)-1]
	h.adjustTopToBottom(0)
	return node
}
func (h *minHeap) Top() *heapNode {
	return h.arr[0]
}

func (h *minHeap) Len() int {
	return len(h.arr)
}

func (h *minHeap) adjustTopToBottom(root int) {
	index := (root+1)*2 - 1
	for index < len(h.arr) {
		if index+1 < len(h.arr) && h.compare(h.arr[index+1], h.arr[index]) {
			index++
		}
		if !h.compare(h.arr[root], h.arr[index]) {
			h.arr[index], h.arr[root] = h.arr[root], h.arr[index]
		} else {
			break
		}
		root = index
		index = (root+1)*2 - 1
	}
}

func (h *minHeap) adjustBottomToTop(index int) {

	//root := (index - 1) / 2
	root := (index+1)/2 - 1

	for root >= 0 {
		if !h.compare(h.arr[root], h.arr[index]) {
			h.arr[index], h.arr[root] = h.arr[root], h.arr[index]
		} else {
			break
		}
		index = root
		root = (index+1)/2 - 1
	}
}

/*
*
根据回忆写的。
1、漏掉 busyHeap.Top().endTime <= meetings[i][0] 相等的判断，写成 <
2、adjustBottomToTop, root=0时会死循环，应该写成root := (index+1)/2 - 1

执行用时分布
48
ms
击败
85.71%
复杂度分析
消耗内存分布
18.44
MB
击败
85.71%
*/
func mostBooked(n int, meetings [][]int) int {

	arr := make([]*heapNode, n)
	for i := 0; i < n; i++ {
		arr[i] = &heapNode{
			id:      i,
			endTime: 0,
			count:   0,
		}
	}
	idleHeap := NewMinHeap(arr, func(root, index *heapNode) bool {
		return root.id < index.id
	})
	busyHeap := NewMinHeap([]*heapNode{}, func(root, index *heapNode) bool {
		return root.endTime < index.endTime || (root.endTime == index.endTime && root.id < index.id)
	})

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	mostId, mostCnt := 0, 0
	for i := 0; i < len(meetings); i++ {
		for busyHeap.Len() > 0 && busyHeap.Top().endTime <= meetings[i][0] {
			idleHeap.Push(busyHeap.Pop())
		}
		var node *heapNode
		if idleHeap.Len() > 0 {
			node = idleHeap.Pop()
			node.endTime = meetings[i][1]
		} else if busyHeap.Len() > 0 {
			node = busyHeap.Pop()
			node.endTime = node.endTime + (meetings[i][1] - meetings[i][0])
		}
		node.count++
		busyHeap.Push(node)
		if node.count > mostCnt {
			mostCnt = node.count
			mostId = node.id
		} else if node.count == mostCnt && node.id < mostId {
			mostId = node.id
		}
	}
	return mostId
}

func TestMostBooked(t *testing.T) {

	fmt.Println(mostBooked(3, [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}))
	// 0 1-20
	// 1 2-10 6-8
	// 2 3-5 4-9(10)
	//
}

/**
解答错误
45 / 81 个通过的测试用例
输入
n =
100
meetings =
[[0,1]]

添加到测试用例
输出
63
预期结果
0
*/
