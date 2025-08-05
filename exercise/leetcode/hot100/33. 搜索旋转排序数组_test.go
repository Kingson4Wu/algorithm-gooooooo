package hot100

import (
	"fmt"
	"testing"
)

/**
个人做法，简单二分查找
看完题解
二分后加上乱序数组和有序数组的判断，通过对比头尾的数值即可

时间
4 ms
击败
50%
内存
2.4 MB
击败
100%

*/

/**

整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。



示例 1：

输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
示例 2：

输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
示例 3：

输入：nums = [1], target = 0
输出：-1


提示：

1 <= nums.length <= 5000
-104 <= nums[i] <= 104
nums 中的每个值都 独一无二
题目数据保证 nums 在预先未知的某个下标上进行了旋转
-104 <= target <= 104
*/

func search2(nums []int, target int) int {

	mid := binarySearch(nums, 0, len(nums)-1, target)
	if nums[mid] == target {
		return mid
	}
	return -1
}

func binarySearch(nums []int, start, end, target int) int {

	if start >= end {
		return start
	}
	mid := (end + start) / 2

	/** 前半段有序 */
	if nums[start] <= nums[mid] {
		if target >= nums[start] && target <= nums[mid] {
			return binarySearch(nums, start, mid, target)
		}
		return binarySearch(nums, mid+1, end, target)
	} else {
		if target >= nums[mid+1] && target <= nums[end] {
			return binarySearch(nums, mid+1, end, target)
		}
		return binarySearch(nums, start, mid, target)
	}
}

func TestSearch(t *testing.T) {
	fmt.Println(search2([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search2([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search2([]int{1}, 0))
	fmt.Println(search2([]int{4, 5, 6, 7, 0, 1, 2}, 5))
	fmt.Println(search2([]int{4, 5, 6, 7, 0, 1, 2}, 8))
	fmt.Println(search2([]int{1}, 1))
	fmt.Println(search2([]int{5, 1, 3}, 1))
}
