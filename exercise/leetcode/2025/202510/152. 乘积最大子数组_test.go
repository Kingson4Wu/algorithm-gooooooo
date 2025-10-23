package _02510

import (
	"fmt"
	"testing"
)

/*
*
根据回忆
两个变量，一个包含上一个的最大，一个包含上一个的最小
*/
func maxProduct(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	maxAns, maxF, minF := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		//记得放同一行，不然重复计算干扰结果
		maxF, minF = max(nums[i]*maxF, max(nums[i]*minF, nums[i])), min(nums[i]*maxF, min(nums[i]*minF, nums[i]))
		maxAns = max(maxAns, maxF)
	}
	return maxAns
}

/**
解答错误
124 / 191 个通过的测试用例

官方题解
输入
nums =
[-4,-3,-2]

添加到测试用例
输出
72
预期结果
12
*/

func TestMaxProduct(t *testing.T) {
	fmt.Println(maxProduct([]int{-4, -3, -2}))
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
5.13
MB
击败
74.38%

*/
