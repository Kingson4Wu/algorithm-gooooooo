package nowcoder

import "math"

/*
*
根据回忆写的
递归中如果发现子树不平衡，直接返回-1，上层全部结束继续递归

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
7.17
MB
击败
18.03%
复杂度分析
*/
func isBalanced(root *TreeNode) bool {

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var height func(root *TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := height(root.Left)
		if left == -1 {
			return -1
		}
		right := height(root.Right)
		if right == -1 {
			return -1
		}
		if math.Abs(float64(left-right)) > 1 {
			return -1
		}
		return 1 + max(left, right)
	}
	return height(root) != -1
}

/**

方法一：自顶向下

class Solution {
public:
    map<TreeNode*, int> hs;
    int depth(TreeNode *root) {
        if (!root) return 0;
        if (hs.find(root) != hs.end()) return hs[root];
        int ldep = depth(root->left);
        int rdep = depth(root->right);
        return hs[root] = max(ldep, rdep) + 1;
    }
    bool judge(TreeNode *root) {
        if (!root) return true;
        return abs(hs[root->left] - hs[root->right]) <= 1 &&
        judge(root->left) && judge(root->right);
    }
    bool IsBalanced_Solution(TreeNode* root) {
        depth(root);
        return judge(root);
    }
};

方法二：自底向上
class Solution {
public:
    int depth(TreeNode *root) {
        if (!root) return 0;
        int ldep = depth(root->left);
        if (ldep == -1) return -1;
        int rdep = depth(root->right);
        if (rdep == -1) return -1;
        int sub = abs(ldep - rdep);
        if (sub > 1) return -1;
        return max(ldep, rdep) + 1;
    }
    bool IsBalanced_Solution(TreeNode* root) {
        return depth(root) != -1;
    }
};
*/
