package hot100

import (
	"fmt"
	"testing"
)

/*
*
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。

输入：n = 3
输出：5
示例 2：

输入：n = 1
输出：1
*/

/*
*
看了题解回忆
关键核心，以不同的元素作为根元素，所组成的树是不同的

我的思路竟然是递归, 估计会超时（竟然没超时）

时间
812 ms
击败
0.86%
内存
1.8 MB
击败
56.78%
*/
func numTrees(n int) int {

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	num := 0
	for i := 0; i < n; i++ {
		if i == 0 || i == n-1 {
			num += numTrees(n - 1)
			continue
		}
		num += numTrees(i) * numTrees(n-i-1)
	}

	return num
}

func TestNumTrees(t *testing.T) {
	fmt.Println(numTrees(1))
	fmt.Println(numTrees(2))
	fmt.Println(numTrees(3))
	fmt.Println(numTrees(4))
}
