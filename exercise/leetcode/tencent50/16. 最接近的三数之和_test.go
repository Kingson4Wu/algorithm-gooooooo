package tencent50

/*
*
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。

示例 1：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
示例 2：

输入：nums = [0,0,0], target = 1
输出：0

提示：

3 <= nums.length <= 1000
-1000 <= nums[i] <= 1000
-104 <= target <= 104
*/
func threeSumClosest(nums []int, target int) int {

	return 0
}

/**

先回去参考三数之和（15. 三数之和）

先排序！

import (
	"sort"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		ans := 0
		for _, l := range nums {
			ans += l
		}
		return ans
	}

	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		start := i + 1
		end := len(nums) - 1
		for start < end {
			tmp := nums[i] + nums[start] + nums[end]
			if math.Abs(float64(tmp-target)) < math.Abs(float64(ans-target)) {
				ans = tmp
			}
			if tmp > target {
				end--
			} else {
				start++
			}
		}
	}

	return ans
}

*/
