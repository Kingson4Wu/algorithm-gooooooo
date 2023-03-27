package hot100

import (
	"fmt"
	"testing"
)

/*
*
自己做的

取决于左边的高度(题目理解错误)

输入
height =
[1,2,1]
添加到测试用例
输出
1
预期结果
2
*/
func maxArea(height []int) int {

	if len(height) <= 1 {
		return 0
	}

	leftMaxHeightIndex := 0
	maxAreaResult := 0

	for i := 1; i < len(height); i++ {
		area := 0
		if height[leftMaxHeightIndex] < height[i] {
			area = height[leftMaxHeightIndex] * (i - leftMaxHeightIndex)
			leftMaxHeightIndex = i
		} else {
			area = height[i] * (i - leftMaxHeightIndex)
		}
		if area > maxAreaResult {
			maxAreaResult = area
		}
	}
	return maxAreaResult
}

func TestMaxArea(t *testing.T) {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
	fmt.Println(maxArea([]int{1}))
}
