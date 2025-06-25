package topinterview150

import (
	"fmt"
	"testing"
)

/*
自己想的
使用数组作为hashmap，下标不超过citations长度

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.09
MB
击败
23.45%

【可以不需要】用负数存为0的情况，取巧

*
给你一个整数数组 citations ，其中 citations[i] 表示研究者的第 i 篇论文被引用的次数。计算并返回该研究者的 h 指数。

根据维基百科上 h 指数的定义：h 代表“高引用次数” ，一名科研人员的 h 指数 是指他（她）至少发表了 h 篇论文，并且 至少 有 h 篇论文被引用次数大于等于 h 。如果 h 有多种可能的值，h 指数 是其中最大的那个。

示例 1：

输入：citations = [3,0,6,1,5]
输出：3
解释：给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 3, 0, 6, 1, 5 次。

	由于研究者有 3 篇论文每篇 至少 被引用了 3 次，其余两篇论文每篇被引用 不多于 3 次，所以她的 h 指数是 3。

示例 2：

输入：citations = [1,3,1]
输出：1
*/
func hIndex(citations []int) int {

	h := 0
	stat := make([]int, len(citations)+1)
	for i := 0; i < len(citations); i++ {
		if citations[i] > len(citations) {
			citations[i] = len(citations)
		}
		if citations[i] == 0 {
			continue
		}
		/*if stat[citations[i]] < 0 {
			stat[citations[i]] = -stat[citations[i]]
		}*/
		stat[citations[i]]++
		for j := 1; j < citations[i]; j++ {
			stat[j]++
			/*if stat[j] > 0 {
				stat[j]++
			} else {
				stat[j]--
			}*/

		}

	}

	for i := 1; i < len(stat); i++ {
		if stat[i] >= i && i > h {
			h = i
		}
	}

	return h
}

func TestHIndex(t *testing.T) {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
	fmt.Println(hIndex([]int{1, 3, 1}))
	fmt.Println(hIndex([]int{1, 0}))
	fmt.Println(hIndex([]int{1, 2, 2}))
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5, 4}))
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5, 4, 7}))

	fmt.Println(hIndex([]int{3, 0, 6, 1, 5, 5, 7}))
}

/**
3
1
1
2
3
4
*/
