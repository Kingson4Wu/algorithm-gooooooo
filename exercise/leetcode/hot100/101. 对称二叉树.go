package hot100

/**
nowcoder/BM31 对称的二叉树.go

自己以前使用栈实现的

现在试试使用递归

时间
0 ms
击败
100%
内存
2.7 MB
击败
99.86%
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	var check func(left, right *TreeNode) bool
	check = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}

		if left != nil && right == nil {
			return false
		}

		if left == nil && right != nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}
		return check(left.Left, right.Right) && check(left.Right, right.Left)
	}

	return check(root.Left, root.Right)

}
