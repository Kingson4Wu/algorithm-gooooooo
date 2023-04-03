package hot100

import (
	"fmt"
	"testing"
)

/**
个人想法
先对角线二分查找，再分别行和列二分查找 （错的，这个规律不对，看看用例中的15，不符合这个规律）
从起始行遍历到最大行，从起始列遍历到最大列

虽然是自己写的，但感觉写得太复杂太乱了

使用官方答案的第三种方法，Z型，从右上角开始比对，移动即可

时间
28 ms
击败
7.45%
内存
6.6 MB
击败
32.2%

官方答案：

方法二：二分查找
思路与算法

由于矩阵 matrix 中每一行的元素都是升序排列的，因此我们可以对每一行都使用一次二分查找，判断 target 是否在该行中，从而判断 target 是否出现。

方法三：Z 字形查找
思路与算法

垃圾说明，直接看代码就好了

我们可以从矩阵 matrix 的右上角 (0,n−1) 进行搜索。在每一步的搜索过程中，如果我们位于位置 (x,y)，那么我们希望在以 matrix 的左下角为左下角、以 (x,y) 为右上角的矩阵中进行搜索，即行的范围为 [x,m−1]，列的范围为 [0,y]：

如果 matrix[x,y]=target，说明搜索完成；

如果 matrix[x,y]>target，由于每一列的元素都是升序排列的，那么在当前的搜索矩阵中，所有位于第 y 列的元素都是严格大于 target 的，因此我们可以将它们全部忽略，即将 y 减少 1；

如果 matrix[x,y]<target，由于每一行的元素都是升序排列的，那么在当前的搜索矩阵中，所有位于第 x 行的元素都是严格小于 target 的，因此我们可以将它们全部忽略，即将 x 增加 1。

在搜索的过程中，如果我们超出了矩阵的边界，那么说明矩阵中不存在 target。

func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    x, y := 0, n-1
    for x < m && y >= 0 {
        if matrix[x][y] == target {
            return true
        }
        if matrix[x][y] > target {
            y--
        } else {
            x++
        }
    }
    return false
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/search-a-2d-matrix-ii/solutions/1062538/sou-suo-er-wei-ju-zhen-ii-by-leetcode-so-9hcx/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

/**
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
输出：true

输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
输出：false


*/

func searchMatrix(matrix [][]int, target int) bool {

	/** 对角线查找 */
	startX, startY := 0, 0
	endX, endY := len(matrix[0])-1, len(matrix)-1

	if matrix[startY][startX] == target {
		return true
	}
	if matrix[endY][endX] == target {
		return true
	}

	for endX-startX > 1 || endY-startY > 1 {

		midX := (startX + endX) / 2
		midY := (startY + endY) / 2

		if matrix[midY][midX] == target {
			return true
		}
		if matrix[midY][midX] > target {
			endX, endY = midX, midY
			continue
		}
		if matrix[midY][midX] < target {
			startX, startY = midX, midY
			continue
		}
	}

	/** 行和列二分查找 */
	x := endX

	for x < len(matrix[0]) {
		startYY, endYY := 0, endY
		for startYY < endYY {

			if endYY-startYY <= 1 {
				if matrix[startYY][x] == target {
					return true
				}
				if matrix[endYY][x] == target {
					return true
				}
				break
			}

			mid := (endYY + startYY) / 2
			if matrix[mid][x] == target {
				return true
			}
			if matrix[mid][x] > target {
				endYY = mid
				continue
			}
			if matrix[mid][x] < target {
				startYY = mid
				continue
			}
		}
		x++
	}

	y := endY
	for y < len(matrix) {
		startXX, endXX := 0, endX
		for startXX < endXX {

			if endXX-startXX <= 1 {
				if matrix[y][startXX] == target {
					return true
				}
				if matrix[y][endXX] == target {
					return true
				}
				break
			}

			mid := (endXX + startXX) / 2
			if matrix[y][mid] == target {
				return true
			}
			if matrix[y][mid] > target {
				endXX = mid
				continue
			}
			if matrix[y][mid] < target {
				startXX = mid
				continue
			}
		}
		y++
	}

	return false
}

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

func checkMatrix(matrix [][]int, t *testing.T) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if !searchMatrix(matrix, matrix[i][j]) {
				t.Fatal("search error,", matrix[i][j])
			}
		}
	}
}
