package dp

import (
	"fmt"
	"testing"
)

/*
*
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

示例 1：

输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
示例 2：

输入：nums = [1,2,3,1]
输出：4
解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。

	偷窃到的最高金额 = 1 + 3 = 4 。

示例 3：

输入：nums = [1,2,3]
输出：3

提示：

1 <= nums.length <= 100
0 <= nums[i] <= 1000

看题解的

dp[i] 表示在下标范围 [start,i] 内可以偷窃到的最高总金额
dp[i]=max(dp[i−2]+nums[i],dp[i−1])

边界条件为：

{
dp[start]=nums[start]
dp[start+1]=max(nums[start],nums[start+1])

只有一间房屋，则偷窃该房屋 ：dp[start]=nums[start]
只有两间房屋，偷窃其中金额较高的房屋：dp[start+1]=max(nums[start],nums[start+1])
分别取 (start,end)=(0,n−2) 和 (start,end)=(1,n−1) 进行计算，取两个 dp[end] 中的最大值，即可得到最终结果

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
3.76
MB
击败
98.93%
*/
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	return max(_rob(nums[:len(nums)-1]), _rob(nums[1:]))
}

func _rob(nums []int) int {

	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func TestRob(t *testing.T) {
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{1, 2, 3}))
	fmt.Println(rob([]int{1, 2}))
	fmt.Println(rob([]int{4}))

}
