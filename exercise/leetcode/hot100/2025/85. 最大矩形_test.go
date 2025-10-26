package _025

import (
	"fmt"
	"testing"
)

/**
单调栈 + 柱状图
将每一行看作是 “柱状图” 的底边（类似于「84. 柱状图中最大矩形」），往上堆积，构建高度数组，然后在每一层使用单调栈求柱状图最大矩形面积。
对每一行 i：
若 matrix[i][j] == '1'，则 heights[j] += 1
否则，heights[j] = 0（清零）


解答错误
11 / 74 个通过的测试用例

官方题解
输入
matrix =
[["1","0"]]

添加到测试用例
输出
0
预期结果
1

执行用时分布
3
ms
击败
61.50%
复杂度分析
消耗内存分布
7.96
MB
击败
40.37%
*/

func maximalRectangle(matrix [][]byte) int {

	heights := make([]int, len(matrix[0]))

	maxArea := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		maxArea = max(largestRectangleArea(heights), maxArea)
	}
	return maxArea
}

func largestRectangleArea(heights []int) int {
	// 在头尾添加两个哨兵，避免特殊情况处理
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	//前哨兵 0	数组头部	保证栈底永不为空，简化边界处理
	//后哨兵 0	数组尾部	强制触发清空栈，确保计算所有可能的最大矩形

	var stack []int // 存的是下标
	maxArea := 0

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			w := i - stack[len(stack)-1] - 1
			// stack[len(stack)-1] 是在出栈之后访问的 → 实际上是原来的 stack[len(stack)-2]。
			//其实算的是上一个的面积
			//确定宽度的边界：
			//1.右边界：右边第一个比它小的柱子就是当前的 i，所以右边界是 i - 1
			//2.左边界：弹出后，新的栈顶 stack[len(stack)-1] 就是左边第一个比它小的柱子，即stack[len(stack)-2]
			maxArea = max(maxArea, h*w)
		}
		stack = append(stack, i)
	}
	return maxArea
}

func TestMaximalRectangle(t *testing.T) {
	fmt.Println(maximalRectangle([][]byte{{'1', '0'}}))
	//fmt.Println(maximalRectangle([][]byte{{'0', '1'}, {'0', '1'}}))

}
