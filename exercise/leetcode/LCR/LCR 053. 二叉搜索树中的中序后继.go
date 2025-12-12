package LCR

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
想半天，感觉只会直接中序遍历

执行用时分布
11
ms
击败
94.34%
复杂度分析
消耗内存分布
9.22
MB
击败
5.66%
*/
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {

	var seekNext bool
	var ans *TreeNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if ans != nil {
			return
		}
		dfs(root.Left)
		if seekNext {
			ans = root
			seekNext = false // 这句不能删除！！！
			return
		}
		if root.Val == p.Val {
			seekNext = true
		}
		dfs(root.Right)
	}
	dfs(root)
	return ans
}

/**
解答错误
13 / 24 个通过的测试用例

官方题解
输入
root =
[5,3,6,2,4,null,null,1]
p =
1

添加到测试用例
输出
5
预期结果
2
生成错误，请检查格式是否正确
*/
