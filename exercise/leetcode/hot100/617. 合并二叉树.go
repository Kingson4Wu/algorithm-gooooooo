package hot100

/**

使用递归轻松写出来了

时间
16 ms
击败
52.13%
内存
6.9 MB
击败
40.30%
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	head := root1
	head.Val = root1.Val + root2.Val
	head.Left = mergeTrees(root1.Left, root2.Left)
	head.Right = mergeTrees(root1.Right, root2.Right)
	return head
}
