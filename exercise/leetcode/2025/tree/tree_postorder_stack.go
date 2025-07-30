package tree

import "algorithm-gooooooo/leetcode"

// 递归写法
/*func postorder(node *leetcode.TreeNode) []int {
	var ans []int
	var dfs func(node *leetcode.TreeNode)
	dfs = func(node *leetcode.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		ans = append(ans, node.Val)
	}
	dfs(node)
	return ans
}*/

// 栈写法
func postorder(node *leetcode.TreeNode) []int {
	var ans []int
	var stack []*leetcode.TreeNode
	//pre用来记录子结点是否已经遍历过的
	var pre *leetcode.TreeNode
	if node != nil {
		stack = append(stack, node)
	}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		if (cur.Left == nil && cur.Right == nil) || (pre != nil && (pre == cur.Left || pre == cur.Right)) {
			ans = append(ans, cur.Val)
			stack = stack[:len(stack)-1]
			pre = cur
		} else {
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
		}
	}
	return ans
}
