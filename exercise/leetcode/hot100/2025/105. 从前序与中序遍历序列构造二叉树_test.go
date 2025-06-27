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
前序：中左右
中序：左中右
以前序第一个为root
通过root节点将中序列表分开左右；再通过中序左边数量，将前序相同数量的找出来
*/

func buildTree(preorder []int, inorder []int) *TreeNode {

	root := &TreeNode{Val: preorder[0]}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			left := inorder[0:i]
			if len(left) > 0 {
				root.Left = buildTree(preorder[1:1+len(left)], left)
			}
			right := inorder[i+1:]
			if len(right) > 0 {
				root.Right = buildTree(preorder[1+len(left):], right)
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
5.54
MB
击败
97.50%

*/
