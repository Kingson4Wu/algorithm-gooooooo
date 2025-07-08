package topinterview150

import (
	"fmt"
	"sort"
	"testing"
)

/**
假设 力扣（LeetCode）即将开始 IPO 。为了以更高的价格将股票卖给风险投资公司，力扣 希望在 IPO 之前开展一些项目以增加其资本。 由于资源有限，它只能在 IPO 之前完成最多 k 个不同的项目。帮助 力扣 设计完成最多 k 个不同项目后得到最大总资本的方式。

给你 n 个项目。对于每个项目 i ，它都有一个纯利润 profits[i] ，和启动该项目需要的最小资本 capital[i] 。

最初，你的资本为 w 。当你完成一个项目时，你将获得纯利润，且利润将被添加到你的总资本中。

总而言之，从给定项目中选择 最多 k 个不同项目的列表，以 最大化最终资本 ，并输出最终可获得的最多资本。

答案保证在 32 位有符号整数范围内。



示例 1：

输入：k = 2, w = 0, profits = [1,2,3], capital = [0,1,1]
输出：4
解释：
由于你的初始资本为 0，你仅可以从 0 号项目开始。
在完成后，你将获得 1 的利润，你的总资本将变为 1。
此时你可以选择开始 1 号或 2 号项目。
由于你最多可以选择两个项目，所以你需要完成 2 号项目以获得最大的资本。
因此，输出最后最大化的资本，为 0 + 1 + 3 = 4。
示例 2：

输入：k = 3, w = 0, profits = [1,2,3], capital = [0,1,2]
输出：6


提示：

1 <= k <= 105
0 <= w <= 109
n == profits.length
n == capital.length
1 <= n <= 105
0 <= profits[i] <= 104
0 <= capital[i] <= 109
*/
/**
根据题解回忆

1. profits和capital成对，按capital升序排序
2. 优先队列按纯收益profits最大堆入列
3. 遍历1的列表，将成本小于等于当前资本的入列，然后pop（资本增加），再继续遍历入列
*/

type IPOHeap struct {
	arr [][]int
}

func (h *IPOHeap) Push(v []int) {
	h.arr = append(h.arr, v)
	h.adjustDownToUp(len(h.arr) - 1)
}

func (h *IPOHeap) adjustUpToDown(root, len int) {
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

func (h *IPOHeap) adjustDownToUp(index int) {
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

func (h *IPOHeap) compare(v1 []int, v2 []int) bool {
	return v1[1] > v2[1]
}

func (h *IPOHeap) Pop() []int {
	v := h.arr[0]
	h.arr[0], h.arr[len(h.arr)-1] = h.arr[len(h.arr)-1], h.arr[0]
	h.arr = h.arr[:len(h.arr)-1]
	h.adjustUpToDown(0, len(h.arr))
	return v
}
func (h *IPOHeap) Len() int {
	return len(h.arr)
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {

	arr := make([][]int, len(capital))
	for i := 0; i < len(capital); i++ {
		arr[i] = []int{capital[i], profits[i]}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0] || (arr[i][0] == arr[j][0] && arr[i][1] > arr[j][1])
	})

	h := IPOHeap{}

	for i := 0; i < len(arr); i++ {
		if arr[i][0] <= w {
			h.Push(arr[i])
		} else {
			if h.Len() > 0 {
				w += h.Pop()[1]
				k--
				i--
			}
		}
		if k <= 0 {
			break
		}
	}
	for h.Len() > 0 && k > 0 {
		w += h.Pop()[1]
		k--
	}
	return w
}

/**
执行用时分布
152
ms
击败
14.47%
复杂度分析
消耗内存分布
21.65
MB
击败
5.26%
*/

func TestFindMaximizedCapital(t *testing.T) {
	fmt.Println(findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1}))
	fmt.Println(findMaximizedCapital(3, 0, []int{1, 2, 3}, []int{0, 1, 2}))
}
