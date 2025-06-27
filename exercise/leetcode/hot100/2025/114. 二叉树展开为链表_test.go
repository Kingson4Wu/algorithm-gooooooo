package _025

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。
*/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		right := root.Right
		root.Right = root.Left
		root.Left = nil
		if right != nil {
			head := root.Right
			for head.Right != nil {
				head = head.Right
			}
			head.Right = right
		}
	}
	flatten(root.Right)
}

/**
执行用时分布
2
ms
击败
1.33%
复杂度分析
消耗内存分布
4.83
MB
击败
10.17%
*/
