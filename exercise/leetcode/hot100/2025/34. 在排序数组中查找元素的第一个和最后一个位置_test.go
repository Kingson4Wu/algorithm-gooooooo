package _025

import (
	"fmt"
	"testing"
)

func searchRange(nums []int, target int) []int {

	seekIndex := func(nums []int, start, end, target int, left bool) int {
		for start <= end {
			mid := (start + end) / 2
			if left {
				if target <= nums[mid] {
					end = mid - 1
				} else {
					start = mid + 1
				}
			} else {
				if target >= nums[mid] {
					start = mid + 1
				} else {
					end = mid - 1
				}
			}
		}
		if left {
			return start
		}
		return end
	}

	leftIndex := seekIndex(nums, 0, len(nums)-1, target, true)
	rightIndex := seekIndex(nums, 0, len(nums)-1, target, false)
	if leftIndex > rightIndex {
		return []int{-1, -1}
	}

	return []int{leftIndex,
		rightIndex}
}

func TestSearchRange(t *testing.T) {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{}, 0))
	fmt.Println(searchRange([]int{1, 2, 3}, 3))
}

/**
[3 4]
[-1 -1]
[-1 -1]
[2 2]
*/

/**

func searchRange(nums []int, target int) []int {
	return []int{findBound(nums, target, true), findBound(nums, target, false)}
}

func findBound(nums []int, target int, isFirst bool) int {
	left, right := 0, len(nums)-1
	bound := -1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			bound = mid
			if isFirst {
				// 继续往左边找
				right = mid - 1
			} else {
				// 继续往右边找
				left = mid + 1
			}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return bound
}

*/
