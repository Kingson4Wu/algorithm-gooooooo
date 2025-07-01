package _025

import (
	"fmt"
	"testing"
)

/*
*
使用Z字型的搜索策略，从右上角开始（左边小下边大）
注意横坐标左移是j--
*/
func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
			continue
		}
		i++
	}
	return false
}

/**
执行用时分布
14
ms
击败
87.09%
复杂度分析
消耗内存分布
8.00
MB
击败
6.58%
*/

func TestSearchMatrix(t *testing.T) {
	//fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
	//fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20))
	//fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 4))

	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	fmt.Println(searchMatrix(matrix, 4))
	checkMatrix(matrix, t)

	matrix2 := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}, {10, 13, 14}, {18, 21, 23}}
	fmt.Println(searchMatrix(matrix2, 4))
	checkMatrix(matrix2, t)

	matrix3 := [][]int{{1}, {2}, {3}, {10}, {18}}
	fmt.Println(searchMatrix(matrix3, 4))
	checkMatrix(matrix3, t)

	matrix4 := [][]int{{1}}
	fmt.Println(searchMatrix(matrix4, 4))
	checkMatrix(matrix4, t)

}

/**
true
true
false
false
*/

func checkMatrix(matrix [][]int, t *testing.T) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if !searchMatrix(matrix, matrix[i][j]) {
				t.Fatal("search error,", matrix[i][j])
			}
		}
	}
}
