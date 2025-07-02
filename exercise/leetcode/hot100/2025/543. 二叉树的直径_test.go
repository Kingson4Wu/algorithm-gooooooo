package _025

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	var deep func(root *TreeNode) int
	deep = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return 1 + max(deep(root.Left), deep(root.Right))
	}

	val := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		//直径定义是：两个节点之间的边数，而你这里的 1 + left + right 是 节点数，比标准答案多了 1。
		v := deep(root.Left) + deep(root.Right)
		if v > val {
			val = v
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	// 感觉这个写法很多重复计算
	return val
}

/**
执行用时分布
579
ms
击败
5.20%
复杂度分析
消耗内存分布
10.28
MB
击败
5.43%
复杂度分析

*/

/**
用后序遍历，在遍历中顺带计算每个节点的深度，并更新最大直径，只遍历一次

func diameterOfBinaryTree(root *TreeNode) int {
    maxDiameter := 0

    var depth func(node *TreeNode) int
    depth = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        left := depth(node.Left)
        right := depth(node.Right)

        // 更新最大直径，边数 = left + right（不加1）
        maxDiameter = max(maxDiameter, left + right)

        return 1 + max(left, right)
    }

    depth(root)
    return maxDiameter
}
*/
