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

/*
*
方法一：优先队列
思路与算法

对于「最大值」，我们可以想到一种非常合适的数据结构，那就是优先队列（堆），其中的大根堆可以帮助我们实时维护一系列元素中的最大值。

作者：力扣官方题解
链接：https://leetcode.cn/problems/sliding-window-maximum/solutions/543426/hua-dong-chuang-kou-zui-da-zhi-by-leetco-ki6m/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

方法二：单调队列
思路与算法

我们可以顺着方法一的思路继续进行优化。

方法三：分块 + 预处理
思路与算法

除了基于「随着窗口的移动实时维护最大值」的方法一以及方法二之外，我们还可以考虑其他有趣的做法。
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
