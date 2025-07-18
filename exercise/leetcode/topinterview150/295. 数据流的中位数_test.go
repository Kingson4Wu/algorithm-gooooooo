package topinterview150

import (
	"fmt"
	"testing"
)

/**
中位数是有序整数列表中的中间值。如果列表的大小是偶数，则没有中间值，中位数是两个中间值的平均值。

例如 arr = [2,3,4] 的中位数是 3 。
例如 arr = [2,3] 的中位数是 (2 + 3) / 2 = 2.5 。
实现 MedianFinder 类:

MedianFinder() 初始化 MedianFinder 对象。

void addNum(int num) 将数据流中的整数 num 添加到数据结构中。

double findMedian() 返回到目前为止所有元素的中位数。与实际答案相差 10-5 以内的答案将被接受。

示例 1：

输入
["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
[[], [1], [2], [], [3], []]
输出
[null, null, null, 1.5, null, 2.0]

解释
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1);    // arr = [1]
medianFinder.addNum(2);    // arr = [1, 2]
medianFinder.findMedian(); // 返回 1.5 ((1 + 2) / 2)
medianFinder.addNum(3);    // arr[1, 2, 3]
medianFinder.findMedian(); // return 2.0
提示:

-105 <= num <= 105
在调用 findMedian 之前，数据结构中至少有一个元素
最多 5 * 104 次调用 addNum 和 findMedian
*/
/**
看完题解回忆
两个优先队列（最大堆和最小堆）
保持两个队列的数量相差不到1
最大堆保存较小的数，最大堆保持较大的数!!!!不要忘了
push时可以将两个堆的数据适当转移，保持数量相差不到1
中位数就是最大堆的顶和最小堆的顶之间的平均数（当总数时偶数时），否则就是数量多的堆顶
adjust ： for root >= 0 { ！！

*/

type heap struct {
	min bool
	arr []int
}

func NewHeap(min bool) *heap {
	return &heap{
		min: min,
	}
}

func (h *heap) adjustUpToDown(root, len int) {
	index := root*2 + 1
	for index < len {
		if index+1 < len && h.compare(h.arr[index+1], h.arr[index]) {
			index++
		}
		if h.compare(h.arr[index], h.arr[root]) {
			h.arr[index], h.arr[root] = h.arr[root], h.arr[index]
		} else {
			break
		}
		root = index
		index = root*2 + 1
	}
}

func (h *heap) adjustDownToUp(index int) {
	root := (index - 1) / 2
	for root >= 0 {
		if h.compare(h.arr[index], h.arr[root]) {
			h.arr[index], h.arr[root] = h.arr[root], h.arr[index]
		} else {
			break
		}
		index = root
		root = (index - 1) / 2
	}
}

func (h *heap) compare(num1, num2 int) bool {
	if h.min {
		return num1 < num2
	} else {
		return num1 > num2
	}
}

func (h *heap) Push(v int) {
	h.arr = append(h.arr, v)
	h.adjustDownToUp(len(h.arr) - 1)
}

func (h *heap) Pop() int {
	v := h.arr[0]
	h.arr[0], h.arr[len(h.arr)-1] = h.arr[len(h.arr)-1], h.arr[0]
	h.arr = h.arr[:len(h.arr)-1]
	h.adjustUpToDown(0, len(h.arr))
	return v
}

func (h *heap) Top() int {
	return h.arr[0]
}

func (h *heap) Len() int {
	return len(h.arr)
}

type MedianFinder struct {
	maxHeap   *heap
	minHeap   *heap
	middleVal float64
}

func Constructor3() MedianFinder {

	return MedianFinder{
		maxHeap: NewHeap(false),
		minHeap: NewHeap(true),
	}
}

func (m *MedianFinder) AddNum(num int) {
	defer m.calMiddle()
	if m.maxHeap.Len() == 0 {
		m.maxHeap.Push(num)
		return
	}
	if float64(num) < m.middleVal {
		m.maxHeap.Push(num)
	} else {
		m.minHeap.Push(num)
	}
	//调整
	for m.minHeap.Len()-m.maxHeap.Len() > 1 {
		m.maxHeap.Push(m.minHeap.Pop())
	}
	for m.maxHeap.Len()-m.minHeap.Len() > 1 {
		m.minHeap.Push(m.maxHeap.Pop())
	}
}

func (m *MedianFinder) calMiddle() {
	if m.maxHeap.Len() == m.minHeap.Len() {
		m.middleVal = float64(m.maxHeap.Top()+m.minHeap.Top()) / 2
	} else if m.maxHeap.Len() > m.minHeap.Len() {
		m.middleVal = float64(m.maxHeap.Top())
	} else {
		m.middleVal = float64(m.minHeap.Top())
	}
}

func (m *MedianFinder) FindMedian() float64 {
	return m.middleVal
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

/**
解答错误
4 / 21 个通过的测试用例

官方题解
输入
["MedianFinder","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian"]
[[],[-1],[],[-2],[],[-3],[],[-4],[],[-5],[]]

添加到测试用例
输出
[null,null,-1.00000,null,-1.50000,null,-2.00000,null,-2.50000,null,-4.00000]
预期结果
[null,null,-1.0,null,-1.5,null,-2.0,null,-2.5,null,-3.0]
*/
/**
执行用时分布
76
ms
击败
83.27%
复杂度分析
消耗内存分布
21.48
MB
击败
68.77%

*/

func TestFindMedian(t *testing.T) {
	m := Constructor3()
	m.AddNum(-1)
	fmt.Println(m.FindMedian())
	m.AddNum(-2)
	fmt.Println(m.FindMedian())
	m.AddNum(-3)
	fmt.Println(m.FindMedian())
	m.AddNum(-4)
	fmt.Println(m.FindMedian())
	m.AddNum(-5)
	fmt.Println(m.FindMedian())
}
