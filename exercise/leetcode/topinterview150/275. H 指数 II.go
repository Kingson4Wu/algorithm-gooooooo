package topinterview150

func hIndex2(citations []int) int {
	h := 0
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] > h {
			h++
		} else {
			break
		}
	}
	return h
}

/**
核心思想
你可以把 h 理解成“目前能确认有多少篇论文的引用数 ≥ h”。
每往左走一篇引用较少的论文时，我们检查它是否仍然能撑起“至少 h+1 篇论文的引用 ≥ h+1”。

这段代码是利用排序后从右向左计数的方式，在每一步都“试探性”地增加 h，
直到遇到引用数不足以支撑下一个 h 的论文，就停下。
此时的 h 就是最大 H 指数。
*/

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
8.88
MB
击败
7.47%
*/

/**
citations =
[0]

添加到测试用例
输出
1
预期结果
0

输入
citations =
[1,1]

添加到测试用例
输出
2
预期结果
1
*/
