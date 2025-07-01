package hot100

import (
	"sort"
	"testing"
)

/**
给定一个会议时间安排的数组，每个会议时间都会包括开始和结束的时间 [[s1,e1],[s2,e2],…] (si < ei)，为避免会议冲突，同时要考虑充分利用会议室资源，请你计算至少需要多少间会议室，才能满足这些会议安排。

示例 1:输入: [[0, 30],[5, 10],[15, 20]]输出: 2
示例 2:输入: [[7,10],[2,4]]输出: 1
2、思路
排序 + 优先队列（最小堆）
	•	将所有会议按照开始时间排序，优先队列存储会议的结束时间由小到大排序
	•	对于当前的会议，如果开始时间小于优先队列元素的会议结束时间，就需要新开一个房间，因此将它重新入队
	•	最后优先队列剩余的就是最少的会议室数

相当于按时间线模拟整个会议的进行

会议开始时间升序排序，然后用一个最小堆（min heap）来保存当前所有正在进行的会议的结束时间。
最小堆代表了所有正在占用的会议室的结束时间。堆顶元素就是最早结束的会议时间

最小堆优势
快速找出最早结束的会议	O(1) 查堆顶
插入新会议结束时间	O(log n)
移除已结束的会议	O(log n)


*/

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// 按开始时间排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 创建一个最小堆，按会议结束时间排序
	roomHeap := &MinHeap{}

	// 第一个会议一定要一个房间，放入堆
	roomHeap.Push(intervals[0][1])

	for i := 1; i < len(intervals); i++ {
		// 如果当前会议的开始时间 >= 最早结束的会议
		if intervals[i][0] >= roomHeap.Top() {
			// 复用该会议室（移除旧会议）
			roomHeap.Pop()
		}
		// 放入当前会议的结束时间
		roomHeap.Push(intervals[i][1])
	}

	// 堆中剩余元素就是所需会议室数
	return roomHeap.Len()
}

type MinHeap struct {
	heap []int
}

func (h *MinHeap) Push(val int) {
	h.heap = append(h.heap, val)
	//Push：只需要自底向上调整（sift up）
	h.siftUp(len(h.heap) - 1)
}

func (h *MinHeap) Pop() int {
	if len(h.heap) == 0 {
		panic("heap is empty")
	}
	top := h.heap[0]
	last := len(h.heap) - 1
	h.heap[0] = h.heap[last]
	h.heap = h.heap[:last]
	//Pop：只需要自顶向下调整（sift down）
	h.siftDown(0)
	return top
}

func (h *MinHeap) Top() int {
	if len(h.heap) == 0 {
		panic("heap is empty")
	}
	return h.heap[0]
}

func (h *MinHeap) Len() int {
	return len(h.heap)
}

func (h *MinHeap) siftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.heap[parent] <= h.heap[i] {
			break
		}
		h.heap[parent], h.heap[i] = h.heap[i], h.heap[parent]
		i = parent
	}
}

func (h *MinHeap) siftDown(i int) {
	n := len(h.heap)
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		if left < n && h.heap[left] < h.heap[smallest] {
			smallest = left
		}
		if right < n && h.heap[right] < h.heap[smallest] {
			smallest = right
		}
		if smallest == i {
			break
		}
		h.heap[i], h.heap[smallest] = h.heap[smallest], h.heap[i]
		i = smallest
	}
}

func TestMinMeetingRooms(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      int
	}{
		{
			intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			want:      2,
		},
		{
			intervals: [][]int{{7, 10}, {2, 4}},
			want:      1,
		},
		{
			intervals: [][]int{{1, 5}, {2, 6}, {3, 7}, {4, 8}},
			want:      4,
		},
		{
			intervals: [][]int{{0, 5}, {5, 10}, {10, 15}},
			want:      1,
		},
		{
			intervals: [][]int{{10, 20}},
			want:      1,
		},
		{
			intervals: [][]int{},
			want:      0,
		},
	}

	for i, tt := range tests {
		got := minMeetingRooms(tt.intervals)
		if got != tt.want {
			t.Errorf("Test case %d failed: expected %d, got %d", i+1, tt.want, got)
		}
	}
}
