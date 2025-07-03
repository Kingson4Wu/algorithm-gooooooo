package hot100

import (
	"fmt"
	"testing"
)

/*
*
个人想法
单调栈（存下标，高度值是单调递增）；另加一个对应的辅助栈存高度
遇到比前面矮的，将高度置为跟矮的一样；相同高度的只保留下标最小的

执行用时分布
439
ms
击败
5.03%
复杂度分析
消耗内存分布
9.27
MB
击败
95.25%
*/
func largestRectangleArea(heights []int) int {

	var indexStack []int
	var heightStack []int

	maxVal := 0
	for i := 0; i < len(heights); i++ {

		if heights[i] == 0 {
			indexStack = nil
			heightStack = nil
			continue
		}

		if len(indexStack) == 0 || heightStack[len(indexStack)-1] < heights[i] {
			indexStack = append(indexStack, i)
			heightStack = append(heightStack, heights[i])
		} else if heightStack[len(indexStack)-1] > heights[i] {

			heightStack[len(indexStack)-1] = heights[i]
			for len(indexStack) > 1 {
				if heightStack[len(heightStack)-1] <= heightStack[len(heightStack)-2] {
					heightStack[len(heightStack)-2] = heightStack[len(heightStack)-1]
					heightStack = heightStack[0 : len(heightStack)-1]
					indexStack = indexStack[0 : len(indexStack)-1]
				} else {
					break
				}
			}
		}
		for k := 0; k < len(indexStack); k++ {
			index := indexStack[k]
			maxVal = max(maxVal, (i-index+1)*heightStack[k])
		}
	}
	return maxVal
}

func TestLargestRectangleArea(t *testing.T) {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{2, 4}))
	fmt.Println(largestRectangleArea([]int{1, 4, 5, 6, 7, 2, 3, 1, 8, 9, 1, 3, 4, 5, 6, 8, 3, 4, 5, 3, 2, 1, 5, 6, 8}))
	fmt.Println(largestRectangleArea([]int{6, 5, 3, 2, 3, 5, 7, 8, 9, 6, 5, 5, 5, 5, 4, 2, 4, 4, 4, 5, 6, 6, 7, 8, 8, 8, 9, 9, 6, 9}))
	fmt.Println(largestRectangleArea([]int{3, 3}))
	fmt.Println(largestRectangleArea([]int{4}))
	fmt.Println(largestRectangleArea([]int{4, 4, 6, 7, 0, 3, 4, 2, 5, 0, 0, 0, 5, 6, 7, 8, 3, 2, 4, 0}))
}

/**
单调栈 + 哨兵法
时间复杂度：O(n)，空间复杂度：O(n)

func largestRectangleArea(heights []int) int {
	// 在头尾添加两个哨兵，避免特殊情况处理
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	//前哨兵 0	数组头部	保证栈底永不为空，简化边界处理
	//后哨兵 0	数组尾部	强制触发清空栈，确保计算所有可能的最大矩形

	stack := []int{} // 存的是下标
	maxArea := 0

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			w := i - stack[len(stack)-1] - 1
			// stack[len(stack)-1] 是在出栈之后访问的 → 实际上是原来的 stack[len(stack)-2]。
			maxArea = max(maxArea, h*w)
		}
		stack = append(stack, i)
	}
	return maxArea
}

用栈保存递增的柱子下标，当遇到更矮的柱子时，开始出栈并计算以弹出柱子为高的最大面积，直到单调性恢复。

优势总结：
只用一个栈，空间少；
不需要每次都遍历整个栈求面积，避免 O(n²)；
哨兵能统一处理边界；
只在单调性被破坏时弹栈，利用“前面的高度更高”的性质求面积，符合人类直觉。

执行用时分布
7
ms
击败
49.96%
复杂度分析
消耗内存分布
9.75
MB
击败
60.67%
*/
