package topinterview145

import (
	"fmt"
	"math"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**

递归超限制
改用队列，层次遍历来实现

理解错题目，不一定是相邻的就一定最小
要用中序遍历来计算！！！！用递归做简单

超出内存限制
最后执行的输入
查看测试用例
root =
[4,2,6,1,3]

*/
/*
*
给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

差值是一个正数，其数值等于两值之差的绝对值。
*/
/*func getMinimumDifference(root *TreeNode) int {

	min := math.MaxInt32
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left != nil {
			if node.Val-node.Left.Val < min {
				min = node.Val - node.Left.Val
			}
			dfs(root.Left)
		}
		if node.Right != nil {
			if node.Right.Val-node.Val < min {
				min = node.Right.Val - node.Val
			}
			dfs(root.Right)
		}
	}
	dfs(root)
	return min
}*/

/*func getMinimumDifference(root *TreeNode) int {

	min := math.MaxInt32
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			if node.Val-node.Left.Val < min {
				min = node.Val - node.Left.Val
			}
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			if node.Right.Val-node.Val < min {
				min = node.Right.Val - node.Val
			}
			queue = append(queue, node.Right)
		}
	}

	return min
}*/

/**
解答错误
129 / 190 个通过的测试用例

官方题解
输入
root =
[236,104,701,null,227,null,911]

添加到测试用例
输出
123
预期结果
9

				236
			/		\
		104			701
			\			\
			227			911

236 - 227 = 9 !!!

*/

func getMinimumDifference(root *TreeNode) int {

	min := math.MaxInt32
	var pre *TreeNode

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left != nil {
			dfs(node.Left)
		}
		if pre != nil {
			if node.Val-pre.Val < min {
				min = node.Val - pre.Val
			}
		}
		pre = node

		if node.Right != nil {
			dfs(node.Right)
		}
	}
	dfs(root)

	return min
}

func TestGetMinimumDifference(t *testing.T) {
	fmt.Println(getMinimumDifference(&TreeNode{Val: 4,
		Right: &TreeNode{Val: 6},
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3}}}))
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
8.45
MB
击败
38.81%
*/
