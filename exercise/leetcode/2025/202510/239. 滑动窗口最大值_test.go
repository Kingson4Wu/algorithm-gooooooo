package _02510

import (
	"fmt"
	"testing"
)

/*
*
根据回忆
用一个递减栈保存当前窗口的后续潜在最大值，第一个是当前窗口最大值，保存的是数组下标（方便后续判断是否在窗口内）
*/
func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, len(nums)-k+1)
	var stack []int

	push := func(i int) {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	//初始入栈, k数量的窗口元素
	for i := 0; i < k; i++ {
		push(i)
	}
	ans[0] = nums[stack[0]]
	for i := 1; i < len(nums)-k+1; i++ {
		index := i + k - 1
		push(index)
		for index-stack[0]+1 > k {
			stack = stack[1:]
		}
		ans[i] = nums[stack[0]]
	}
	return ans
}

func TestMaxSlidingWindow(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{7, 2, 4}, 2))
}

/**
7
ms
击败
74.17%
复杂度分析
消耗内存分布
9.98
MB
击败
75.93%

*/
