package _025

import (
	"fmt"
	"testing"
)

/*
*
自己想的
就是栈 存升序的下标， 配合maxVal，endIndex变量，得到最终的子序列下标
寻找规律
*/
func findUnsortedSubarray(nums []int) int {

	// 存升序的下标
	var stack []int
	endIndex := 0
	stack = append(stack, 0)
	maxVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] >= maxVal {
			maxVal = nums[i]
			if endIndex == 0 {
				stack = append(stack, i)
			}
		} else {
			for len(stack) > 0 && nums[i] < nums[stack[len(stack)-1]] {
				stack = stack[0 : len(stack)-1]
			}
			endIndex = i
		}
	}
	var startIndex int
	if len(stack) == 0 {
		startIndex = 0
	} else {
		startIndex = stack[len(stack)-1] + 1
	}
	if len(stack) == len(nums) {
		return 0
	}
	return endIndex - startIndex + 1
}

/**
解答错误
277 / 307 个通过的测试用例

官方题解
输入
nums =
[2,3,3,2,4]

添加到测试用例
输出
2
预期结果
3

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
8.27
MB
击败
41.42%
*/

func TestFindUnsortedSubarray(t *testing.T) {

	fmt.Println(findUnsortedSubarray([]int{2, 3, 3, 2, 4}))

	fmt.Println(findUnsortedSubarray([]int{1, 3, 5, 4, 3, 4, 5, 6, 7, 4, 5, 6, 9}))

	fmt.Println(findUnsortedSubarray([]int{5, 3, 4, 1, 4, 6, 7, 2, 4, 7, 8, 9}))

	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4, 7, 4, 3, 5, 2, 8, 9}))

	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4}))
	fmt.Println(findUnsortedSubarray([]int{1}))
}

/**
3
10
9
7
5
0
0
*/
