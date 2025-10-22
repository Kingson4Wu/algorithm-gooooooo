package _02510

import (
	"fmt"
	"testing"
)

func maxArea(height []int) int {
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	left, right := 0, len(height)-1
	ans := 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if area > ans {
			ans = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
9.48
MB
击败
48.26%
*/

func TestMaxArea(t *testing.T) {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

}
