package hot100

import (
	"fmt"
	"testing"
)

/*
*

给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

示例 1：

输入：nums = [1,5,11,5]
输出：true
解释：数组可以分割成 [1, 5, 5] 和 [11] 。
示例 2：

输入：nums = [1,2,3,5]
输出：false
解释：数组不能分割成两个元素和相等的子集。
*/

/*
好难想的动态规划

target = sum/2

dp[i][target] 表示前i个（0...i）是否能选中若干个的和是等于target
1、nums[i] > target; 不能选； dp[i][target] = dp[i-1][target]
2、nums[i] <= target; 选或者不选；dp[i][target] = dp[i-1][target-num[i]] ||  dp[i-1][target]

初始化：
1、dp[i][0] = true 都不选
2、dp[0][nums[0]] = true

边界条件考虑，提前结束
1、数组总和sum和为奇数
2、最大的数字 > sum/2
*/
/**
时间
28 ms
击败
46.27%
内存
6.9 MB
击败
36.71%
*/
func canPartition(nums []int) bool {

	if len(nums) < 2 {
		return false
	}

	sum := 0
	maxNum := 0
	for _, num := range nums {
		sum += num
		if maxNum < num {
			maxNum = num
		}
	}
	//和为奇数
	if sum%2 != 0 {
		return false
	}
	if maxNum > sum/2 {
		return false
	}

	dp := make([][]bool, len(nums))
	target := sum / 2
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, target+1)
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true

	/** dp[0][nums[i]] 这些可以不用填充  */

	for i := 1; i < len(dp); i++ {
		for j := 1; j <= target; j++ {
			if nums[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j-nums[i]] || dp[i-1][j]
			}
		}
	}

	for i := 0; i < len(dp); i++ {
		if dp[i][target] {
			return true
		}
	}
	return false
}

func TestCanPartition(t *testing.T) {

	fmt.Println(canPartition([]int{1, 5, 11, 5}))
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
}
