package _02510

import (
	"fmt"
	"testing"
)

/*
*
还是重写看答案才会！！！
*/
func searchRange(nums []int, target int) []int {

	seekIndex := func(nums []int, target int, left bool) int {
		lower, upper := 0, len(nums)-1
		bound := -1
		for lower <= upper {
			mid := (lower + upper) / 2
			if nums[mid] == target {
				bound = mid
				if left {
					upper = mid - 1
				} else {
					lower = mid + 1
				}
			} else if nums[mid] < target {
				lower = mid + 1
			} else {
				upper = mid - 1
			}
		}
		return bound
	}
	return []int{seekIndex(nums, target, true), seekIndex(nums, target, false)}
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
