package _025

import "algorithm-gooooooo/leetcode"

// 递归写法
/*func inorder(node *leetcode.TreeNode) []int {
	var ans []int
	var dfs func(node *leetcode.TreeNode)
	dfs = func(node *leetcode.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		ans = append(ans, node.Val)
		dfs(node.Right)
	}
	dfs(node)
	return ans
}*/

// 栈写法
func inorder(node *leetcode.TreeNode) []int {
	var ans []int
	var stack []*leetcode.TreeNode
	for len(stack) > 0 || node != nil {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, node.Val)
			node = node.Right
		}
	}
	return ans
}

// 判断是否是二叉搜索数
func isValidBST(node *leetcode.TreeNode) bool {
	var pre *leetcode.TreeNode
	result := true
	var dfs func(node *leetcode.TreeNode)
	dfs = func(node *leetcode.TreeNode) {
		if !result {
			return
		}
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != nil && pre.Val > node.Val {
			result = false
			return
		}
		pre = node
		dfs(node.Right)
	}
	dfs(node)
	return result
}
