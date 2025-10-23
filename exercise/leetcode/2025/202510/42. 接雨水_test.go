package _02510

import (
	"fmt"
	"testing"
)

// 左右最高的
/*func trap(height []int) int {

	leftHeights := make([]int, len(height))
	rightHeights := make([]int, len(height))

	for i := 0; i < len(height); i++ {
		if i == 0 {
			leftHeights[i] = height[i]
			continue
		}
		if height[i] > leftHeights[i-1] {
			leftHeights[i] = height[i]
		} else {
			leftHeights[i] = leftHeights[i-1]
		}
	}
	for i := len(height) - 1; i >= 0; i-- {
		if i == len(height)-1 {
			rightHeights[i] = height[i]
			continue
		}
		if height[i] > rightHeights[i+1] {
			rightHeights[i] = height[i]
		} else {
			rightHeights[i] = rightHeights[i+1]
		}
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	area := 0
	for i := 0; i < len(height); i++ {
		area += min(leftHeights[i], rightHeights[i]) - height[i]
	}
	return area
}*/

func trap(height []int) (ans int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])

		if leftMax < rightMax {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
7.80
MB
击败
45.73%
*/

func TestTrap(t *testing.T) {
	fmt.Println(trap([]int{10, 1, 2, 9}))
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
	fmt.Println(trap([]int{4, 2, 3}))
	fmt.Println(trap([]int{4, 2}))
}

/**
6
9
1
0
*/

/**

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
7.78
MB
击败
61.44%

*/
