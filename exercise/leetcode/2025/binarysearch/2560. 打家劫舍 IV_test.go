package binarysearch

import (
	"fmt"
	"testing"
)

/**
自己按回忆写的
middle := (low + high) / 2 可能不会nums中，但是符合条件会一直缩小范围，知道刚好不符合跳出循环，此时的lower就是在nums中
*/

// 另给你一个整数 k ，表示窃贼将会窃取的 最少 房屋数。小偷总能窃取至少 k 间房屋。
// 返回小偷的 最小 窃取能力。
func minCapability(nums []int, k int) int {

	low, high := nums[0], nums[0]

	for _, num := range nums {
		if num < low {
			low = num
		}
		if num > high {
			high = num
		}
	}

	for low <= high {
		middle := (low + high) / 2
		visited := false
		count := 0

		for _, num := range nums {
			if !visited && num <= middle {
				count++
				visited = true
			} else {
				visited = false
			}
			if count >= k {
				break
			}
		}
		if count >= k {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return low
}

func TestMinCapability(t *testing.T) {
	//fmt.Println(minCapability([]int{2, 3, 5, 9}, 2))
	fmt.Println(minCapability([]int{2, 7, 9, 3, 1}, 2))
	fmt.Println(minCapability([]int{3, 6, 8, 10}, 2))
}
