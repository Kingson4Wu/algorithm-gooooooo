package hot100

import (
	"fmt"
	"testing"
)

/**
不会做
看题解

三种方法
1、动态规划，分别两趟算出下标为i的左边和右边的最高点，使用两个list存，最后一趟算出每个i的容量
2、单调栈（从左到右）
3、双指针（左右开始往中间）


时间
8 ms
击败
91.94%
内存
5.9 MB
击败
38.61%
*/

/*
*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：

输入：height = [4,2,0,3,2,5]
输出：9

提示：

n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105
*/
func trap(height []int) int {

	if len(height) < 3 {
		return 0
	}

	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	leftMax[0] = height[0]
	rightMax[len(height)-1] = height[len(height)-1]
	for i := 1; i < len(height); i++ {
		if height[i] > leftMax[i-1] {
			leftMax[i] = height[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	}
	for i := len(height) - 2; i >= 0; i-- {
		if height[i] > rightMax[i+1] {
			rightMax[i] = height[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}

	area := 0
	for i := 1; i < len(height)-1; i++ {
		area += min(leftMax[i], rightMax[i]) - height[i]
	}

	return area
}

func TestTrap(t *testing.T) {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
	fmt.Println(trap([]int{4, 2, 3}))
	fmt.Println(trap([]int{4, 2}))
}
