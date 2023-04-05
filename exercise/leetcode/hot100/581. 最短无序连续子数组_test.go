package hot100

import (
	"fmt"
	"math"
	"testing"
)

/**
方法1：先排序、在对比
这里使用方法2：
找出中间段的左右边界
先从左到右找中间段的右下标（即遍历过程中最后一个小于左边最大值的）right
再找出从右到左中间段的左下标（即遍历过程中最后一个大于右边最小值的）left
right-left +1 即结果

时间
20 ms
击败
87.76%
内存
6.3 MB
击败
62.15%

*/

/**
自己总结的过滤好像有点小复杂
已经不知道在写啥了，也不过。。。
*/

/*
*
给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。

请你找出符合题意的 最短 子数组，并输出它的长度。

示例 1：

输入：nums = [2,6,4,8,10,9,15]
输出：5
解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
示例 2：

输入：nums = [1,2,3,4]
输出：0
示例 3：

输入：nums = [1]
输出：0
*/
func findUnsortedSubarray(nums []int) int {

	left, right := -1, -1
	leftMin, rightMax := math.MaxInt, math.MinInt

	for i := 0; i < len(nums); i++ {
		if nums[i] >= rightMax {
			rightMax = nums[i]
		} else {
			right = i
		}
		if nums[len(nums)-1-i] <= leftMin {
			leftMin = nums[len(nums)-1-i]
		} else {
			left = len(nums) - 1 - i
		}
	}

	if right == -1 {
		return 0
	}
	return right - left + 1
}

/*func findUnsortedSubarray(nums []int) int {

	if len(nums) <= 1 {
		return 0
	}

	start, end := -1, -1
	leftMax, max := nums[0], nums[0]

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if start == -1 {
				start = i
				end = i + 1
				max = nums[i]
				continue
			}
			end = i + 1
		}
		if start == -1 {
			continue
		}
		if nums[i+1] < max {
			end = i + 1
		}
		if nums[i+1] <= leftMax {
			for start-1 >= 0 {
				if nums[start-1] > nums[i+1] {
					start--
					leftMax = nums[start]
				} else {
					break
				}
			}
		}
	}

	if end > start {

		for start-1 >= 0 {
			if nums[start-1] == nums[start] {
				start--
			} else {
				break
			}
		}

		return end - start + 1
	}

	return 0
}*/

func TestFindUnsortedSubarray(t *testing.T) {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4}))
	fmt.Println(findUnsortedSubarray([]int{1}))
	fmt.Println(findUnsortedSubarray([]int{1, 3, 2, 3, 3}))
	fmt.Println(findUnsortedSubarray([]int{2, 3, 3, 2, 4}))

	fmt.Println(findUnsortedSubarray([]int{2, 3, 3, 2, 4, 5, 3}))
	fmt.Println(findUnsortedSubarray([]int{2, 3, 3, 2, 4, 5, 1}))
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 3, 2, 4, 5, 1}))

	fmt.Println(findUnsortedSubarray([]int{1, 1, 2, 3, 3, 2, 4, 5, 1}))

	fmt.Println(findUnsortedSubarray([]int{1, 1, 2, 3, 3, 2, 4, 5, 1, 0}))

	fmt.Println(findUnsortedSubarray([]int{5, 4, 3, 10, 1, 8, 9, 11, 12, 13}))
	/**
	输入
	nums =
	[2,3,3,2,4]
	添加到测试用例
	输出
	2
	预期结果
	3
	*/

}
