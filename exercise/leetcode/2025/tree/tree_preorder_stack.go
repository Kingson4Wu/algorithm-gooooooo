package tree

import "algorithm-gooooooo/leetcode"

// 递归写法
/*func preorder(node *leetcode.TreeNode) []int {
	var ans []int
	var dfs func(node *leetcode.TreeNode)
	dfs = func(node *leetcode.TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(node)
	return ans
}*/

// 栈写法
func preorder(node *leetcode.TreeNode) []int {
	var ans []int
	var stack []*leetcode.TreeNode
	for len(stack) > 0 || node != nil {
		if node != nil {
			ans = append(ans, node.Val)
			stack = append(stack, node)
			node = node.Left
		} else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node = node.Right
		}
	}
	return ans
}
