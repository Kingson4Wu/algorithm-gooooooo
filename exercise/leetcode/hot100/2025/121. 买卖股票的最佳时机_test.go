package _025

import (
	"fmt"
	"testing"
)

/*
*
这么简单，为啥之前想那么复杂

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
10.18
MB
击败
13.76%
复杂度分析
*/
func maxProfit(prices []int) int {

	profit := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > profit {
			profit = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return profit
}

func TestMaxProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
