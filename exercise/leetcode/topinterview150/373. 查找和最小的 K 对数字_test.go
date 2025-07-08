package topinterview150

import (
	"fmt"
	"testing"
)

/**
给定两个以 非递减顺序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。

定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。

请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。



示例 1:

输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
输出: [1,2],[1,4],[1,6]
解释: 返回序列中的前 3 对数：
     [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
示例 2:

输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
输出: [1,1],[1,1]
解释: 返回序列中的前 2 对数：
     [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]


提示:

1 <= nums1.length, nums2.length <= 105
-109 <= nums1[i], nums2[i] <= 109
nums1 和 nums2 均为 升序排列
1 <= k <= 104
k <= nums1.length * nums2.length
*/
/**
感觉双指针就能搞定
不是双指针这么简单

---

用优先队列
m = len(num1)
n = len(num2)

入列规则：
1. 一开始将num1的所有入列，num2只取下标0即可
则需要入num1的min(k,m)的长度即可，因为如果k>m, num1的下标已经全入列了，如果k<=m, 答案也只会在num1的k个下标内
2. 出列在入列，入列取的元素就是出列数据i,j 入一个j+1即可，因为下一个只有可能在堆里或j+1

*/

type PairHeap struct {
	arr   [][]int
	nums1 []int
	nums2 []int
}

func (h *PairHeap) Push(v []int) {
	h.arr = append(h.arr, v)
	h.adjustDownToUp(len(h.arr) - 1)
}

func (h *PairHeap) adjustUpToDown(root, len int) {
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

func (h *PairHeap) adjustDownToUp(index int) {
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

func (h *PairHeap) compare(v1 []int, v2 []int) bool {
	return h.nums1[v1[0]]+h.nums2[v1[1]] < h.nums1[v2[0]]+h.nums2[v2[1]]
}

func (h *PairHeap) Pop() []int {
	v := h.arr[0]
	h.arr[0], h.arr[len(h.arr)-1] = h.arr[len(h.arr)-1], h.arr[0]
	h.arr = h.arr[:len(h.arr)-1]
	h.adjustUpToDown(0, len(h.arr))
	return v
}
func (h *PairHeap) Len() int {
	return len(h.arr)
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

	var ans [][]int
	m := len(nums1)
	n := len(nums2)

	heap := PairHeap{
		nums1: nums1,
		nums2: nums2,
	}
	for i := 0; i < m && i < k; i++ {
		heap.Push([]int{i, 0})
	}
	for len(ans) < k {
		v := heap.Pop()
		ans = append(ans, []int{nums1[v[0]], nums2[v[1]]})
		//这个最小的还有下一个，否则不用再push了，剩下的结果都在堆里
		if v[1]+1 < n {
			heap.Push([]int{v[0], v[1] + 1})
		}
	}
	return ans
}

/**
执行用时分布
28
ms
击败
51.52%
复杂度分析
消耗内存分布
10.86
MB
击败
73.83%
复杂度分析

*/

func TestKSmallestPairs(t *testing.T) {
	fmt.Println(kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))
	fmt.Println(kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 9))
}
