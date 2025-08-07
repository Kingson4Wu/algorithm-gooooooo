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
