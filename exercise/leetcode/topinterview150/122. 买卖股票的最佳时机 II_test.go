package topinterview150

/*
*
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.89
MB
击败
98.64%
*/
/**
因为交易次数不受限，如果可以把所有的升序的全部收集到，一定是利益最大化的
*/
func maxProfit(prices []int) int {
	var ans int
	if len(prices) <= 1 {
		return 0
	}
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}
