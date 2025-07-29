package tree

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	//找到要删除的结点
	if root.Left == nil {
		return root.Right
	}
	if root.Right == nil {
		return root.Left
	}
	// 情况4：左右子树都有
	//需要找右子树的最小值的结点, 最小的一定在最左边
	successor := root.Right
	for successor.Left != nil {
		successor = successor.Left
	}
	successor.Right = deleteNode(root.Right, successor.Val)
	successor.Left = root.Left
	return successor
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
8.91
MB
击败
54.96%

*/

/**
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的key对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/delete-node-in-a-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
作者：LeetCode-Solution
链接：https://leetcode.cn/problems/delete-node-in-a-bst/solution/shan-chu-er-cha-sou-suo-shu-zhong-de-jie-n6vo/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/
