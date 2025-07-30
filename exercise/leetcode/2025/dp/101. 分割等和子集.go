package dp

// 给定一个非空的正整数数组 nums ，请判断能否将这些数字分成元素和相等的两部分。
// 简单题？！竟然不会做。。
// 动态规划
// 状态定义：dp[i][j] 表示前 i 个数中，是否存在子集和为 j。
/**
凭记忆写的
执行用时分布
19
ms
击败
42.31%
复杂度分析
消耗内存分布
8.33
MB
击败
38.46%
*/
func canPartition(nums []int) bool {

	if len(nums) == 1 {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2

	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] > target {
			return false
		}
		if nums[i] == target {
			return true
		}
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}
	dp[0][0] = true
	dp[0][nums[0]] = true

	for i := 1; i < len(nums); i++ {
		for j := 1; j <= target; j++ {
			if nums[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[len(nums)-1][target]
}
