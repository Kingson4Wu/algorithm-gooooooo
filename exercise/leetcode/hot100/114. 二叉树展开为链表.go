package hot100

/**
看了题解，原地算法（O(1) 额外空间），关键是寻找树的最右结点
注意left要置空
把左边结点放右边，右边结点放在原左结点的最后一个最结点，然后一直递归右子树，最子树置空

时间
0 ms
击败
100%
内存
2.8 MB
击败
13.96%

题解的方法1是先序遍历，直接新建另一棵树
方法2则是原地算法
*/

/**
给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。

进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {

	if root == nil {
		return
	}
	if root.Left == nil {
		flatten(root.Right)
		return
	}
	if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		flatten(root.Right)
		return
	}
	rightNode := traceLastRightNode(root.Left)
	rightNode.Right = root.Right
	root.Right = root.Left
	root.Left = nil
	flatten(root.Right)
}

func traceLastRightNode(root *TreeNode) *TreeNode {

	if root.Left == nil && root.Right == nil {
		return root
	}
	if root.Left == nil {
		return traceLastRightNode(root.Right)
	}
	if root.Right == nil {
		return traceLastRightNode(root.Left)
	}
	return traceLastRightNode(root.Right)
}
