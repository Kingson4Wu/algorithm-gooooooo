package hot100

import (
	"fmt"
	"testing"
)

/**
自己做的，拆分子问题，通过递归可解,但是超时

超出时间限制
41 / 63 个通过的测试用例
最后执行的输入
添加到测试用例
m =
51
n =
9

todo 估计要改成栈使用回溯


*/

/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？
输入：m = 3, n = 7
输出：28
示例 2：

输入：m = 3, n = 2
输出：3
解释：
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右
3. 向下 -> 向右 -> 向下
示例 3：

输入：m = 7, n = 3
输出：28
示例 4：

输入：m = 3, n = 3
输出：6
*/

func uniquePaths(m int, n int) int {

	if m == 0 || n == 0 {
		return 0
	}

	if m == 1 || n == 1 {
		return 1
	}

	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
}

func TestUniquePaths(t *testing.T) {
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(7, 3))
	fmt.Println(uniquePaths(3, 3))
	fmt.Println(uniquePaths(5, 5))
}
