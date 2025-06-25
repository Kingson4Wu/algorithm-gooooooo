package topinterview150

import (
	"fmt"
	"testing"
)

/**
好像自己分析完就写完了
用滑动窗口维护一个处于中间的可交换区间即可，符合条件的（不超过2的）交换到前面来

执行用时分布
6
ms
击败
75.20%
复杂度分析
消耗内存分布
6.87
MB
击败
53.49%

*/
/*
*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

示例 1：

输入：nums = [1,1,1,2,2,3]
输出：5, nums = [1,1,2,2,3]
解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3。 不需要考虑数组中超出新长度后面的元素。
示例 2：

输入：nums = [0,0,1,1,1,1,2,3,3]
输出：7, nums = [0,0,1,1,2,3,3]
解释：函数应返回新长度 length = 7, 并且原数组的前七个元素被修改为 0, 0, 1, 1, 2, 3, 3。不需要考虑数组中超出新长度后面的元素。

提示：

1 <= nums.length <= 3 * 104
-104 <= nums[i] <= 104
nums 已按升序排列
*/
func removeDuplicates(nums []int) int {

	start, length := 0, 0
	currentNum, currentCount := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == currentNum {
			if currentCount < 2 {
				currentCount++
				if length > 0 {
					nums[start], nums[i] = nums[i], nums[start]
					start++
				}
			} else {
				if length == 0 {
					start = i
					length = 1
				} else {
					length++
				}
			}
		} else {
			currentNum = nums[i]
			currentCount = 1
			if length > 0 {
				nums[start], nums[i] = nums[i], nums[start]
				start++
			}
		}

	}

	/*fmt.Println("===")
	for i := 0; i < len(nums)-length; i++ {
		fmt.Println(nums[i])
	}
	fmt.Println("===")*/

	return len(nums) - length
}

func TestRemoveDuplicates(t *testing.T) {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))

	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))

	fmt.Println(removeDuplicates([]int{0, 1, 2}))

	fmt.Println(removeDuplicates([]int{0, 1}))

	fmt.Println(removeDuplicates([]int{0, 1, 1, 1}))

	fmt.Println(removeDuplicates([]int{0, 0, 0, 1, 2, 2, 2, 3, 3, 3, 3, 3, 4, 5, 6, 6, 6, 7, 7, 7}))
}
