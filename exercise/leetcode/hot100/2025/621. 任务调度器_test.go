package _025

import (
	"fmt"
	"testing"
)

/*
*
最大堆, 保存所有的任务，按任务数量排序
执行后放入N大的数组，后续执行任务时 任务轮次%n 得到对应的下标（冷却够时间的任务），先入堆

（其实可以不用最大堆，因为任务数量是一直变化的，可以遍历的过程中，维护冷却状态等）
（不过我这个方法目前看比较清晰）


执行用时分布
17
ms
击败
32.68%
复杂度分析
消耗内存分布
8.04
MB
击败
15.03%
*/

type taskNode struct {
	name byte
	num  int
}

type maxHeap struct {
	tasks []*taskNode
}

func newHeap(tasks []*taskNode) *maxHeap {
	heap := &maxHeap{
		tasks: tasks,
	}
	for i := len(tasks)/2 - 1; i >= 0; i-- {
		heap.adjustUp(i, len(tasks))
	}
	return heap
}

func (h *maxHeap) Push(node *taskNode) {
	h.tasks = append(h.tasks, node)
	h.adjustDown(len(h.tasks)-1, len(h.tasks))
}

func (h *maxHeap) Len() int {
	return len(h.tasks)
}

func (h *maxHeap) Pop() *taskNode {
	node := h.tasks[0]
	h.tasks[0], h.tasks[len(h.tasks)-1] = h.tasks[len(h.tasks)-1], h.tasks[0]
	h.tasks = h.tasks[0 : len(h.tasks)-1]
	h.adjustUp(0, len(h.tasks))
	return node
}
func (h *maxHeap) Top() *taskNode {
	return h.tasks[0]
}

func (h *maxHeap) adjustUp(root int, n int) {

	index := root*2 + 1
	for index < n {
		if index+1 < n && h.tasks[index+1].num > h.tasks[index].num {
			index++
		}
		if h.tasks[index].num > h.tasks[root].num {
			h.tasks[root], h.tasks[index] = h.tasks[index], h.tasks[root]
		} else {
			break
		}
		root = index
		index = root*2 + 1
	}
}

func (h *maxHeap) adjustDown(index int, n int) {

	root := (index - 1) / 2
	for root >= 0 {
		if h.tasks[index].num > h.tasks[root].num {
			h.tasks[root], h.tasks[index] = h.tasks[index], h.tasks[root]
		} else {
			break
		}
		index = root
		root = (index - 1) / 2
	}
}

func leastInterval(tasks []byte, n int) int {

	m := make(map[byte]int)
	for _, task := range tasks {
		m[task]++
	}
	taskList := make([]*taskNode, len(m))
	i := 0
	for k, v := range m {
		taskList[i] = &taskNode{name: k, num: v}
		i++
	}
	coldTasks := make([]*taskNode, n+1)
	coldTaskCount := 0

	heap := newHeap(taskList)

	count := 0
	for heap.Len() > 0 || coldTaskCount > 0 {

		index := count % (n + 1)
		if coldTaskCount > 0 {
			if coldTasks[index] != nil {
				heap.Push(coldTasks[index])
				coldTasks[index] = nil
				coldTaskCount--
			}
		}
		if heap.Len() > 0 {
			task := heap.Pop()
			task.num--
			if task.num > 0 {
				coldTasks[index] = task
				coldTaskCount++
			}
		}
		count++
	}
	return count
}

func TestLeastInterval(t *testing.T) {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
	fmt.Println(leastInterval([]byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1))
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 3))

	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0))
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2))
}
