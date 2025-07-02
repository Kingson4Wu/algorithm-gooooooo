package hot100

import (
	"fmt"
	"testing"
)

/**
使用递归解题
没有考虑到0的情况！

时间
252 ms
击败
34.59%
内存
2 MB
击败
81.21%
*/

/*
*

给你一个整数数组 nums 和一个整数 target 。

向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：

例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

示例 1：

输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
示例 2：

输入：nums = [1], target = 1
输出：1
*/
func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 1 {
		/** 0的情况 */
		if nums[0] == 0 && target == 0 {
			return 2
		}
		if nums[0] == target {
			return 1
		}
		if nums[0] == -target {
			return 1
		}

		return 0
	}
	return findTargetSumWays(nums[1:], target-nums[0]) + findTargetSumWays(nums[1:], target+nums[0])
}

func TestFindTargetSumWays(t *testing.T) {
	fmt.Println(findTargetSumWays([]int{1, 1}, 0))
	fmt.Println(findTargetSumWays([]int{1, 1, 1}, 1))
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWays([]int{1}, 1))

	fmt.Println(findTargetSumWays([]int{1, 0}, 1))
	/**
	输入
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
}
