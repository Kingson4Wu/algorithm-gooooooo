package sort

import (
	"fmt"
	"sort"
	"testing"
)

/*
*
（找出被安排次数最多的会议室编号）

所有会议按开始时间排序

构造两个最小堆：

空闲会议室（按编号）
占用会议室（按结束时间 + 编号）
每轮会议安排时：

如果无空闲会议室，则延期
记录每个会议室的使用次数
最终返回使用次数最多的会议室中编号最小者
*/

type meetingNode struct {
	id      int
	count   int
	endTime int
}

type meetingHeap struct {
	arr     []*meetingNode
	compare func(i, j *meetingNode) bool
}

func NewMeetingHead(compare func(i, j *meetingNode) bool) *meetingHeap {
	return &meetingHeap{
		compare: compare,
	}
}

func (h *meetingHeap) Push(node *meetingNode) {
	h.arr = append(h.arr, node)
	h.AdjustFromBottomToTop(h.Len() - 1)
}
func (h *meetingHeap) Pop() *meetingNode {
	node := h.arr[0]
	h.arr[0] = h.arr[h.Len()-1]
	h.arr = h.arr[:h.Len()-1]
	h.AdjustFromTopToBottom(0)
	return node
}

func (h *meetingHeap) Top() *meetingNode {
	return h.arr[0]
}

func (h *meetingHeap) Len() int {
	return len(h.arr)
}

func (h *meetingHeap) AdjustFromTopToBottom(root int) {
	index := root*2 + 1
	for index < h.Len() {
		if index+1 < h.Len() && h.compare(h.arr[index], h.arr[index+1]) {
			index++
		}
		if !h.compare(h.arr[root], h.arr[index]) {
			break
		}
		h.arr[root], h.arr[index] = h.arr[index], h.arr[root]
		root = index
		index = root*2 + 1
	}
}

func (h *meetingHeap) AdjustFromBottomToTop(index int) {
	root := (index+1)/2 - 1
	for root >= 0 {
		if !h.compare(h.arr[root], h.arr[index]) {
			break
		}
		h.arr[root], h.arr[index] = h.arr[index], h.arr[root]
		index = root
		root = (index+1)/2 - 1
	}
}

func mostBooked(n int, meetings [][]int) int {

	if n == 1 {
		return 0
	}

	maxCount := 0
	minId := 0

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	idleMeetingHeap := NewMeetingHead(func(i, j *meetingNode) bool {
		return i.id > j.id
	})
	for i := 0; i < n; i++ {
		idleMeetingHeap.Push(&meetingNode{id: i})
	}
	usedMeetingHeap := NewMeetingHead(func(i, j *meetingNode) bool {
		return i.endTime > j.endTime || (i.endTime == j.endTime && i.id > j.id)
	})
	for _, meeting := range meetings {
		//检查有没有空闲出来的会议室
		//其实一样跟之前一个堆的做法一样，也是要全部检查再push进去
		for usedMeetingHeap.Len() > 0 && usedMeetingHeap.Top().endTime <= meeting[0] {
			idleMeetingHeap.Push(usedMeetingHeap.Pop())
		}
		var room *meetingNode
		//找一个空闲的会议室
		if idleMeetingHeap.Len() > 0 {
			room = idleMeetingHeap.Pop()
			room.endTime = meeting[1]
		} else {
			// 从使用中的会议室找一个最早结束的
			room = usedMeetingHeap.Pop()
			room.endTime += meeting[1] - meeting[0]
		}
		room.count++
		usedMeetingHeap.Push(room)
		if room.count > maxCount || (room.count == maxCount && room.id < minId) {
			minId = room.id
			maxCount = room.count
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

/**
执行用时分布
38
ms
击败
100.00%
复杂度分析
消耗内存分布
17.74
MB
击败
100.00%
复杂度分析

*/

/**
1
0
0
0
1
*/
