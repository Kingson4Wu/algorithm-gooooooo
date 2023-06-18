package hot100

import (
	"fmt"
	"testing"
)

/*
*
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

示例 1：

输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
示例 2：

输入：nums = [0,1,0,3,2,3]
输出：4
示例 3：

输入：nums = [7,7,7,7,7,7,7]
输出：1

进阶：

你能将算法的时间复杂度降低到 O(n log(n)) 吗?
*/

/*
*
方法一：动态规划
思路与算法

定义 dp[i] 为考虑前 i 个元素，以第 i 个数字结尾的最长上升子序列的长度，注意 nums[i] 必须被选取。

我们从小到大计算 dp 数组的值，在计算 dp[i] 之前，我们已经计算出 dp[0…i−1] 的值，则状态转移方程为：
dp[i]=max(dp[j])+1,其中0≤j<i且num[j]<num[i]

作者：力扣官方题解
链接：https://leetcode.cn/problems/longest-increasing-subsequence/solutions/147667/zui-chang-shang-sheng-zi-xu-lie-by-leetcode-soluti/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func lengthOfLIS(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	maxNum := 1

	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {

			dp[i] = max(dp[i], dp[j]+1)
		}
		maxNum = max(maxNum, dp[i])
	}

	return maxNum
}

func TestLengthOfLIS(t *testing.T) {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}
