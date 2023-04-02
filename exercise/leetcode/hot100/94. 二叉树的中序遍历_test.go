package hot100

import (
	"fmt"
	"testing"
)

/**
nowcoder/BM24 二叉树的中序遍历.go

以前做过，使用栈做的

这次使用递归试试

时间
0 ms
击败
100%
内存
1.9 MB
击败
5.58%

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) (ans []int) {

	if root == nil {
		return
	}

	ans = append(ans, inorderTraversal(root.Left)...)
	ans = append(ans, root.Val)
	ans = append(ans, inorderTraversal(root.Right)...)

	return
}

func TestInorderTraversal(t *testing.T) {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}

	fmt.Println(inorderTraversal(root))
}
