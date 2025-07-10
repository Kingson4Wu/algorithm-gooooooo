package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

/**
给你一个整数 n ，共有编号从 0 到 n - 1 的 n 个会议室。

给你一个二维整数数组 meetings ，其中 meetings[i] = [starti, endi] 表示一场会议将会在 半闭 时间区间 [starti, endi) 举办。所有 starti 的值 互不相同 。

会议将会按以下方式分配给会议室：

每场会议都会在未占用且编号 最小 的会议室举办。
如果没有可用的会议室，会议将会延期，直到存在空闲的会议室。延期会议的持续时间和原会议持续时间 相同 。
当会议室处于未占用状态时，将会优先提供给原 开始 时间更早的会议。
返回举办最多次会议的房间 编号 。如果存在多个房间满足此条件，则返回编号 最小 的房间。

半闭区间 [a, b) 是 a 和 b 之间的区间，包括 a 但 不包括 b 。



示例 1：

输入：n = 2, meetings = [[0,10],[1,5],[2,7],[3,4]]
输出：0
解释：
- 在时间 0 ，两个会议室都未占用，第一场会议在会议室 0 举办。
- 在时间 1 ，只有会议室 1 未占用，第二场会议在会议室 1 举办。
- 在时间 2 ，两个会议室都被占用，第三场会议延期举办。
- 在时间 3 ，两个会议室都被占用，第四场会议延期举办。
- 在时间 5 ，会议室 1 的会议结束。第三场会议在会议室 1 举办，时间周期为 [5,10) 。
- 在时间 10 ，两个会议室的会议都结束。第四场会议在会议室 0 举办，时间周期为 [10,11) 。
会议室 0 和会议室 1 都举办了 2 场会议，所以返回 0 。
示例 2：

输入：n = 3, meetings = [[1,20],[2,10],[3,5],[4,9],[6,8]]
输出：1
解释：
- 在时间 1 ，所有三个会议室都未占用，第一场会议在会议室 0 举办。
- 在时间 2 ，会议室 1 和 2 未占用，第二场会议在会议室 1 举办。
- 在时间 3 ，只有会议室 2 未占用，第三场会议在会议室 2 举办。
- 在时间 4 ，所有三个会议室都被占用，第四场会议延期举办。
- 在时间 5 ，会议室 2 的会议结束。第四场会议在会议室 2 举办，时间周期为 [5,10) 。
- 在时间 6 ，所有三个会议室都被占用，第五场会议延期举办。
- 在时间 10 ，会议室 1 和 2 的会议结束。第五场会议在会议室 1 举办，时间周期为 [10,12) 。
会议室 1 和会议室 2 都举办了 2 场会议，所以返回 1 。


提示：

1 <= n <= 100
1 <= meetings.length <= 105
meetings[i].length == 2
0 <= starti < endi <= 5 * 105
starti 的所有值 互不相同
*/

/**
感觉构建一个最小堆就好了

当会议室处于未占用状态时，将会优先提供给原 开始 时间更早的会议。 (不需要提前开！！！)

    - 所有会议先按开始时间排序
    - 以结束时间最小(结束时间一样，按id最小)构建n大的最小堆（n个会议室）（将第一个会议的开始时间作为全部会议室的结束时间）
    - 每次pop之后，当endTime < 将进行会议的 startTime， 则 endTime = startTime，重新push进去；重新pop，知道 endTime >= 将进行会议的 startTime（全部会议室都按此操作一遍）
    - 那么新的endTime = endTime + (会议endTime - 会议startTime)， 重新push
    - 遍历过程中计算count最大对应的最小的id


自己独立做的
执行用时分布
45
ms
击败
100.00%
复杂度分析
消耗内存分布
19.36
MB
击败
11.11%
*/

type minHeap struct {
	heap []*heapNode
}

type heapNode struct {
	id      int
	endTime int
	count   int
}

func newHeap(heap []*heapNode) *minHeap {

	h := &minHeap{
		heap: heap,
	}
	for i := len(heap)/2 - 1; i >= 0; i-- {
		h.adjustUpToDown(i, len(heap))
	}
	return h
}

func compare(node1, node2 *heapNode) bool {
	if node1.endTime < node2.endTime {
		return true
	}
	if node1.endTime == node2.endTime && node1.id < node2.id {
		return true
	}
	return false
}

func (h *minHeap) adjustUpToDown(root, n int) {

	index := root*2 + 1
	for index < n {

		if index+1 < n && compare(h.heap[index+1], h.heap[index]) {
			index++
		}
		if compare(h.heap[index], h.heap[root]) {
			h.heap[index], h.heap[root] = h.heap[root], h.heap[index]
		} else {
			break
		}
		root = index
		index = root*2 + 1
	}
}

func (h *minHeap) adjustDownToUp(index, n int) {

	root := (index - 1) / 2
	for root >= 0 {
		if compare(h.heap[index], h.heap[root]) {
			h.heap[index], h.heap[root] = h.heap[root], h.heap[index]
		} else {
			break
		}
		index = root
		root = (index - 1) / 2
	}
}

// Push 和 Pop 是特殊写法，用下标，而且逻辑是pop之后一定会push
func (h *minHeap) Push(node *heapNode) {
	h.heap[len(h.heap)-1] = node
	h.adjustDownToUp(len(h.heap)-1, len(h.heap))
}

func (h *minHeap) Pop() *heapNode {
	node := h.heap[0]
	h.heap[0], h.heap[len(h.heap)-1] = h.heap[len(h.heap)-1], h.heap[0]
	h.adjustUpToDown(0, len(h.heap)-1)
	return node
}

func mostBooked(n int, meetings [][]int) int {

	if n == 1 {
		return 0
	}

	// 按开始时间排序
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	list := make([]*heapNode, n)
	for i := 0; i < n; i++ {
		list[i] = &heapNode{
			id: i,
			//取最早开始的会议为结束时间，这样较小的id排前面，也减少后面的频繁pop和push（为了修改endTime）
			endTime: meetings[0][0],
			count:   0,
		}
	}
	maxCount, minId := 0, 0
	heap := newHeap(list)

	for i := 0; i < len(meetings); i++ {
		room := heap.Pop()
		// 会议室空闲了，但会议室时间没到
		for room.endTime < meetings[i][0] {
			room.endTime = meetings[i][0]
			heap.Push(room)
			room = heap.Pop()
		}

		room.count++
		room.endTime += meetings[i][1] - meetings[i][0]
		heap.Push(room)

		if room.count > maxCount {
			minId = room.id
			maxCount = room.count
		} else if room.count == maxCount && room.id < minId {
			minId = room.id
		}
	}
	return minId
}

func TestMostBooked(t *testing.T) {

	fmt.Println(mostBooked(2, [][]int{{0, 10}, {8, 12}, {14, 18}, {19, 21}, {20, 31}, {21, 39}, {25, 27}, {28, 30}, {31, 35}}))
	fmt.Println(mostBooked(1, [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}))
	fmt.Println(mostBooked(6, [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}))
	fmt.Println(mostBooked(2, [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}))
	fmt.Println(mostBooked(3, [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}))
}
