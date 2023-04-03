package hot100

/**
nowcoder/BM26 求二叉树的层序遍历.go

以前使用栈做的，这次尝试使用递归

看了递归，实现也不简单，也要辅助hash
但比起要自己维护栈还是简单点

时间
4 ms
击败
34.30%
内存
3 MB
击败
5.75%

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	highMap := make(map[int][]int)
	var highF func(*TreeNode, int)
	highF = func(root *TreeNode, high int) {
		if root == nil {
			return
		}
		highMap[high] = append(highMap[high], root.Val)
		highF(root.Left, high+1)
		highF(root.Right, high+1)
	}

	highF(root, 0)

	var result [][]int
	for i := 0; i < len(highMap); i++ {
		result = append(result, highMap[i])
	}

	return result
}
