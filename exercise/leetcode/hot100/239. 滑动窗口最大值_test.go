package hot100

import (
	"fmt"
	"testing"
)

/**
动态规划可解，但是又超出时间限制！！
*/
/*
*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。

示例 1：

输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3

	1 [3  -1  -3] 5  3  6  7       3
	1  3 [-1  -3  5] 3  6  7       5
	1  3  -1 [-3  5  3] 6  7       5
	1  3  -1  -3 [5  3  6] 7       6
	1  3  -1  -3  5 [3  6  7]      7

示例 2：

输入：nums = [1], k = 1
输出：[1]
*/
func maxSlidingWindow(nums []int, k int) []int {

	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = nums[i]
	}

	for length := 2; length <= k; length++ {
		for i := 0; i < len(nums)-k+1; i++ {
			dp[i] = max(dp[i], nums[i+length-1])
		}
	}

	return dp[:len(nums)-k+1]
}

func TestMaxSlidingWindow(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1}, 1))
}
