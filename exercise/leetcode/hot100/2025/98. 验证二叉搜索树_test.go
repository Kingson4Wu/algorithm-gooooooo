package _025

/*type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
*/
/**
个人想法，用中序遍历，检查
*/
func isValidBST(root *TreeNode) bool {

	var pre *TreeNode
	result := true

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if !result {
			return
		}
		if node.Left != nil {
			dfs(node.Left)
		}
		if pre != nil {
			//if node.Val < pre.Val {
			if node.Val <= pre.Val {
				result = false
				return
			}
		}
		pre = node
		if node.Right != nil {
			dfs(node.Right)
		}
	}
	if root != nil {
		dfs(root)
	}
	return result
}

/**
解答错误
73 / 86 个通过的测试用例

官方题解
输入
root =
[2,2,2]

添加到测试用例
输出
true
预期结果
false
*/

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
7.13
MB
击败
13.64%
*/
