package top100liked

import (
	"fmt"
	"testing"
)

/**

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。


使用二分查找
*/

func searchInsert(nums []int, target int) int {

	start, end := 0, len(nums)-1

	for start <= end {
		if start == end {
			if nums[start] >= target {
				return start
			} else {
				return start + 1
			}
		}
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid
		}
	}
	return 0
}

/**
自己做的

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.57
MB
击败
5.45%
*/

func TestSearchInsert(t *testing.T) {

	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))

	fmt.Println(searchInsert([]int{1, 3, 5, 6, 8}, 5))

	fmt.Println(searchInsert([]int{1, 3, 5, 6, 8}, 3))
	fmt.Println(searchInsert([]int{1, 3, 5, 6, 8}, 4))
	fmt.Println(searchInsert([]int{1, 3, 5, 6, 8}, 9))

	fmt.Println(searchInsert([]int{2}, 9))
	fmt.Println(searchInsert([]int{2}, 1))
	fmt.Println(searchInsert([]int{2}, 2))

}
