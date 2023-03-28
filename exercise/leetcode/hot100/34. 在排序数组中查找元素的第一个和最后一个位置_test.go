package hot100

import (
	"fmt"
	"testing"
)

/**
个人做法：
二分查找

时间
8 ms
击败
51.76%
内存
3.8 MB
击败
26.17%

*/
/**
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。



示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]


提示：

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109
*/
/**
输入
nums =
[1,2,3]
target =
3
添加到测试用例
输出
[0,0]
预期结果
[2,2]
*/
func searchRange(nums []int, target int) []int {

	index := searchIndex(nums, 0, len(nums), target)

	if index == -1 {
		return []int{-1, -1}
	}

	left, right := index, index
	for left-1 >= 0 && nums[left-1] == nums[index] {
		left--
	}
	for right+1 < len(nums) && nums[right+1] == nums[index] {
		right++
	}

	return []int{left, right}
}

func searchIndex(nums []int, start, end, target int) int {
	if len(nums[start:end]) == 0 {
		return -1
	}
	mid := len(nums[start:end])/2 + start

	if nums[mid] > target {
		return searchIndex(nums, start, mid, target)
	} else if nums[mid] < target {
		return searchIndex(nums, mid+1, end, target)
	} else {
		return mid
	}
}

func TestSearchRange(t *testing.T) {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{}, 0))
	fmt.Println(searchRange([]int{1, 2, 3}, 3))
}
