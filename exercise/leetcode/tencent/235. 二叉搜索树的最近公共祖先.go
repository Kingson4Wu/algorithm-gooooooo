package tencent

/**
自己做的
时间
16 ms
击败
82.37%
内存
6.8 MB
击败
93.86%
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	if p.Val >= root.Val && q.Val <= root.Val {
		return root
	}

	if p.Val <= root.Val && q.Val >= root.Val {
		return root
	}

	if p.Val >= root.Val && q.Val >= root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return lowestCommonAncestor(root.Left, p, q)
}
