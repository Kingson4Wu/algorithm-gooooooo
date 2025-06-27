package topinterview150

/**
给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。


*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
*
通过自己画图计算可知，套路和105. 从前序与中序遍历序列构造二叉树 一样
root的在前序的第一个，在后序的最后一个
*/
func buildTree(inorder []int, postorder []int) *TreeNode {

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {

			left := inorder[0:i]
			if len(left) > 0 {
				root.Left = buildTree(left, postorder[0:len(left)])
			}
			right := inorder[i+1:]
			if len(right) > 0 {
				root.Right = buildTree(right, postorder[len(left):len(left)+len(right)])
			}
			break
		}
	}
	return root
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
5.53
MB
击败
96.40%
*/
