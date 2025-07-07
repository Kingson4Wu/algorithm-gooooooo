package topinterview150

import (
	"fmt"
	"testing"
)

/**
给定一个长度为 n 的环形整数数组 nums ，返回 nums 的非空 子数组 的最大可能和 。

环形数组 意味着数组的末端将会与开头相连呈环状。形式上， nums[i] 的下一个元素是 nums[(i + 1) % n] ， nums[i] 的前一个元素是 nums[(i - 1 + n) % n] 。

子数组 最多只能包含固定缓冲区 nums 中的每个元素一次。形式上，对于子数组 nums[i], nums[i + 1], ..., nums[j] ，不存在 i <= k1, k2 <= j 其中 k1 % n == k2 % n 。



示例 1：

输入：nums = [1,-2,3,-2]
输出：3
解释：从子数组 [3] 得到最大和 3
示例 2：

输入：nums = [5,-3,5]
输出：10
解释：从子数组 [5,5] 得到最大和 5 + 5 = 10
示例 3：

输入：nums = [3,-2,2,-3]
输出：3
解释：从子数组 [3] 和 [3,-2,2] 都可以得到最大和 3


提示：

n == nums.length
1 <= n <= 3 * 104
-3 * 104 <= nums[i] <= 3 * 104
*/
/**
hot100/53. 最大子数组和_test.go 无环的
*/

/*
*
想不出来，看题解

情况1：最大子数组不跨环，即最大子数组和_test.go 无环的
情况2：最大子数组跨了环（首尾拼接）
  - 整个数组总和 - 最小连续子数组和 = 最大跨环子数组和

终极策略总结：
1. 先用 Kadane 算法求出最大子数组和：maxSubSum
2. 再用 Kadane 算法求出最小子数组和：minSubSum
3. 计算整个数组总和 totalSum
4. 答案是：
  - 如果 maxSubSum < 0：说明全是负数，只能取 maxSubSum
  - 否则，取 max(maxSubSum, totalSum - minSubSum)
*/
func maxSubarraySumCircular(nums []int) int {

	maxVal := nums[0]
	minVal := nums[0]
	sum := nums[0]
	preMaxSum := nums[0]
	preMinSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if preMaxSum > 0 {
			preMaxSum += nums[i]
		} else {
			preMaxSum = nums[i]
		}
		if preMaxSum > maxVal {
			maxVal = preMaxSum
		}

		if preMinSum < 0 {
			preMinSum += nums[i]
		} else {
			preMinSum = nums[i]
		}
		if preMinSum < minVal {
			minVal = preMinSum
		}

		sum += nums[i]
	}
	if maxVal < 0 {
		return maxVal
	}
	if sum-minVal > maxVal {
		return sum - minVal
	}
	return maxVal
}

func TestMaxSubarraySumCircular(t *testing.T) {
	fmt.Println(maxSubarraySumCircular([]int{1, -2, 3, -2}))
	fmt.Println(maxSubarraySumCircular([]int{5, -3, 5}))
}

/**
执行用时分布
1
ms
击败
37.99%
复杂度分析
消耗内存分布
8.72
MB
击败
35.05%
复杂度分析

*/
