package hot100

import (
	"fmt"
	"sort"
	"testing"
)

/**
看了提示自己做的

没有去重！！
*/

/**
1、先排序，从小到大
2、以三元组第二个数为基准，左右两边双指针往中间计算
3、注意计算结果不重复，需对一些条件跳过当前循环
*/

/**
输入
nums =
[-1,0,1,0]
添加到测试用例
输出
[[-1,0,1],[-1,0,1]]
预期结果
[[-1,0,1]]
*/

/**
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。
示例 2：

输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。
示例 3：

输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0 。


提示：

3 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	var result [][]int

	/** 取三元组第二个，然后往两边寻找 */
	for i := 1; i < len(nums)-1; i++ {

		left, right := 0, len(nums)-1

		/** 相同的数字，已经计算过 */
		if i > 1 && nums[i] == nums[i-1] && nums[i-1] == nums[i-2] {
			continue
		}

		for left < i && right > i {

			/** 防左右重复 */
			if nums[left] == nums[left+1] && left+1 != i {
				left++
				continue
			}
			if nums[right] == nums[right-1] && right-1 != i {
				right--
				continue
			}

			sum := nums[left] + nums[i] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[left], nums[i], nums[right]})
				left++
				right--
				continue
			}
			if sum > 0 {
				right--
				continue
			}
			if sum < 0 {
				left++
				continue
			}
		}

	}

	return result
}

func TestThreeSum(t *testing.T) {

	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, -1, -1, -1}))
	fmt.Println(threeSum([]int{-1, 0, 1, 0}))

}
