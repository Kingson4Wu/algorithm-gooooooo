package hot100

import "testing"

/**
按照之前的印象，尝试完成
1、先序的第一位即树的根结点
2、从中序中找出根结点的位置，除根结点分成两半，同时得到左树的长度
3、得到左子树长度后，也可以把先序除根结点外分成左右两半
4、递归构造左右子树
*/
/**
给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

preorder 和 inorder 均 无重复 元素
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	mid := -1
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			mid = i
			break
		}
	}

	if mid != -1 {
		if mid-1 >= 0 {
			root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
		}
		if mid+1 < len(inorder) {
			root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
		}
	}

	return root
}

/*
*
输入
preorder =
[3,9,20,15,7]
inorder =
[9,3,15,20,7]
输出
[3,null,20,null,7]
预期结果
[3,9,20,null,null,15,7]
*/
func TestBuildTree(t *testing.T) {
	buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
}
