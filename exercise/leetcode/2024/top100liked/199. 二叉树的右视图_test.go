package top100liked

import "testing"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
*/

/**
做错一次，重做
用队列好做，写起来麻烦，怎么用递归呢？没想到， 官方答案也不是递归
最终用队列做，但感觉不够简洁

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
3.95
MB
击败
13.30%
*/

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	var ans []int

	for len(queue) > 0 {
		ans = append(ans, queue[len(queue)-1].Val)
		var newQueue []*TreeNode
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				newQueue = append(newQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				newQueue = append(newQueue, queue[i].Right)
			}
		}
		queue = newQueue
	}
	return ans
}

//=============

/*
*
自己做的，直接用递归方便

解答错误
158 / 216 个通过的测试用例

官方题解
输入
root =
[1,2,3,4]

添加到测试用例
输出
[1,3]
预期结果
[1,3,4]

理解错误，还是得层级遍历
*/
/*func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Right != nil {
		subSideView := rightSideView(root.Right)
		return append([]int{root.Val}, subSideView...)
	}
	if root.Left != nil {
		subSideView := rightSideView(root.Left)
		return append([]int{root.Val}, subSideView...)
	}
	return []int{root.Val}
}*/

func TestRightSideView(t *testing.T) {

}
