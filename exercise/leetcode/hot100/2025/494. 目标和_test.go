package _025

import (
	"fmt"
	"testing"
)

func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 1 {

		// 这种写法没考虑 +0 ，-0的情况 ！！！
		/*if nums[0] == target || -nums[0] == target {
			return 1
		}*/
		count := 0
		if nums[0] == target {
			count++
		}
		if -nums[0] == target {
			count++
		}
		return count
	}
	sub1 := findTargetSumWays(nums[1:], target-nums[0])
	sub2 := findTargetSumWays(nums[1:], target+nums[0])
	return sub1 + sub2
}

/**
执行用时分布
247
ms
击败
24.33%
复杂度分析
消耗内存分布
3.96
MB
击败
88.33%
复杂度分析

*/

func TestFindTargetSumWays(t *testing.T) {
	fmt.Println(findTargetSumWays([]int{1, 1}, 0))
	fmt.Println(findTargetSumWays([]int{1, 1, 1}, 1))
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWays([]int{1}, 1))

	//todo
	fmt.Println(findTargetSumWays([]int{1, 0}, 1))
}

/**
nums =
[1,0]
target =
1

添加到测试用例
输出
1
预期结果
2
*/
