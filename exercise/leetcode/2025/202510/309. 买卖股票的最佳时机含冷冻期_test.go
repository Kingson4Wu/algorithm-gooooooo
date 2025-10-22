package _02510

/*
*
第i天
f[i][0]持有股票
f[i][1]卖出股票（处于冷冻期）
f[i][2]未持有股票（非冷冻期）

f[i][0] = max(f[i-1][0], f[i-1][2]-prices[i])
f[i][1] = f[i-1][0]+prices[i]
f[i][2] = max(f[i-1][1], f[i-1][2])

公式只考虑i-1的就好，不用考虑i-2的

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.00
MB
击败
90.61%
*/
func maxProfit(prices []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	f1, f2, f3 := -prices[0], 0, 0
	for i := 1; i < len(prices); i++ {
		f1, f2, f3 = max(f1, f3-prices[i]), f1+prices[i], max(f2, f3)
	}
	return max(f2, f3)
}
