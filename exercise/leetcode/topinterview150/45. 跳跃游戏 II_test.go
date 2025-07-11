package topinterview150

/*
*
动态规划
*/
func jump(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+nums[i] && j < len(nums); j++ {
			if dp[j] == 0 {
				dp[j] = dp[i] + 1
			} else {
				dp[j] = min(dp[i]+1, dp[j])
			}
		}
	}
	return dp[len(nums)-1]
}

/**
解答错误
63 / 110 个通过的测试用例

官方题解
输入
nums =
[2,1]

添加到测试用例
输出
0
预期结果
1
*/

/**
执行用时分布
27
ms
击败
14.96%
复杂度分析
消耗内存分布
7.66
MB
击败
18.70%
*/
