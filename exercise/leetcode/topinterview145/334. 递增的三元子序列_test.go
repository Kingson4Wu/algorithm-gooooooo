package topinterview145

import "math"

/**
给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。

如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。



示例 1：

输入：nums = [1,2,3,4,5]
输出：true
解释：任何 i < j < k 的三元组都满足题意
示例 2：

输入：nums = [5,4,3,2,1]
输出：false
解释：不存在满足题意的三元组
示例 3：

输入：nums = [2,1,5,0,4,6]
输出：true
解释：三元组 (3, 4, 5) 满足题意，因为 nums[3] == 0 < nums[4] == 4 < nums[5] == 6


提示：

1 <= nums.length <= 5 * 105
-231 <= nums[i] <= 231 - 1


进阶：你能实现时间复杂度为 O(n) ，空间复杂度为 O(1) 的解决方案吗？
*/

/*
*
看题解

直接看答案吧
基本完全看答案写的

方法二：贪心
可以使用贪心的方法将空间复杂度降到 O(1)。从左到右遍历数组 nums，遍历过程中维护两个变量 first 和 second，分别表示递增的三元子序列中的第一个数和第二个数，任何时候都有 first<second。

初始时，first=nums[0]，second=+∞。对于 1≤i<n，当遍历到下标 i 时，令 num=nums[i]，进行如下操作：

如果 num>second，则找到了一个递增的三元子序列，返回 true；

否则，如果 num>first，则将 second 的值更新为 num；

否则，将 first 的值更新为 num。

如果遍历结束时没有找到递增的三元子序列，返回 false。
*/
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	first, second := nums[0], math.MaxInt32
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num > second {
			return true
		} else if num > first {
			second = num
		} else {
			first = num
		}
	}
	return false
}

/**
执行用时分布
4
ms
击败
37.13%
复杂度分析
消耗内存分布
18.68
MB
击败
73.76%
*/
