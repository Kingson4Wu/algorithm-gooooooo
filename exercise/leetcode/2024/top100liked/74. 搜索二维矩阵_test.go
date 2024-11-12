package top100liked

import (
	"fmt"
	"testing"
)

/*
*
给你一个满足下述两条属性的 m x n 整数矩阵：

每行中的整数从左到右按非严格递增顺序排列。
每行的第一个整数大于前一行的最后一个整数。
给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。
*/

/*
*
个人想法，二维转一维，二分查找
*/
/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.17
MB
击败
6.18%
*/
func searchMatrix(matrix [][]int, target int) bool {

	if len(matrix) == 0 {
		return false
	}
	width := len(matrix[0])

	start, end := 0, len(matrix)*len(matrix[0])-1

	// 这里注意要相等！！！
	for start <= end {
		mid := (start + end) / 2
		x := mid % width
		y := mid / width
		// 记住y放第一个下标！！！
		if matrix[y][x] == target {
			return true
		} else if matrix[y][x] > target {
			//这里记住要减1 ！！！
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false
}

func TestSearchMatrix(t *testing.T) {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
}

/**
方法一：两次二分查找
思路

由于每行的第一个元素大于前一行的最后一个元素，且每行元素是升序的，所以每行的第一个元素大于前一行的第一个元素，因此矩阵第一列的元素是升序的。

我们可以对矩阵的第一列的元素二分查找，找到最后一个不大于目标值的元素，然后在该元素所在行中二分查找目标值是否存在。



方法二：一次二分查找
思路

若将矩阵每一行拼接在上一行的末尾，则会得到一个升序数组，我们可以在该数组上二分找到目标元素。

代码实现时，可以二分升序数组的下标，将其映射到原矩阵的行和列上。



结语
两种方法殊途同归，都利用了二分查找，在二维矩阵上寻找目标值。值得注意的是，若二维数组中的一维数组的元素个数不一，方法二将会失效，而方法一则能正确处理。
！！！！！
*/
