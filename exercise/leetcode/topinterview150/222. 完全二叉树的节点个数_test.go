package topinterview150

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://cloud.tencent.com/developer/article/1880865
//普通二叉树写法：时间复杂度 O(N)
/*func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}*/
//满二叉树，节点总数就和树的高度呈指数关系，时间复杂度 O(logN)
/*func countNodes(root *TreeNode) int {
	h := 0
	// 计算树的高度
	for root != nil {
		root = root.Left
		h++
	}
	// 节点总数就是 2^h - 1
	return int(math.Pow(2, float64(h))) - 1
}*/
//完全二叉树
//两个递归只有一个会真的递归下去，另一个一定会触发hl == hr而立即返回，不会递归下去。
//这种写法最后一层都要检查一遍，没有利用到完全二叉树的特性
func countNodes(root *TreeNode) int {
	l, r := root, root
	hl, hr := 0, 0
	for l != nil {
		l = l.Left
		hl++
	}
	for r != nil { // 算到最右的高度
		r = r.Right
		hr++
	}
	// 如果左右子树的高度相同，则是一棵满二叉树
	if hl == hr {
		return int(math.Pow(2, float64(hl))) - 1
	}
	// 如果左右高度不同，则按照普通二叉树的逻辑计算
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

//官网对最后一层进行二分查找
/**
					  1
				/        	\
 			 2          	    3
          / 	\             /   \
 		4        5           6      7
       / \      / \        / \      / \
      8   9   10  (11)  (12) (13) (14) (15)
假设括号的不存在
二分查找区间[8,15]
int low = 1 << level, high = (1 << (level + 1)) - 1;
int mid = (high - low + 1) / 2 + low;
(high + low + 1) /2 = 12
查找12是否存在 = 1100
int bits = 1 << (level - 1);  = 100
bits & k   --- 100 & 1100 = 100 > 0  (要找12，从1结点走右子树，即到3)
bits >>= 1;   = 10
bits & k   --- 10 & 1100 = 0 == 0  (要找12，从3结点走左子树，即到6)
bits >>= 1;   = 1
bits & k   --- 1 & 1100 = 0 == 0  (要找12，从6结点走左子树，即到12)

*/

/**
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    level := 0
    for node := root; node.Left != nil; node = node.Left {
        level++
    }
    return sort.Search(1<<(level+1), func(k int) bool {
        if k <= 1<<level {
            return false
        }
        bits := 1 << (level - 1)
        node := root
        for node != nil && bits > 0 {
            if bits&k == 0 {
                node = node.Left
            } else {
                node = node.Right
            }
            bits >>= 1
        }
        return node == nil
    }) - 1
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/count-complete-tree-nodes/solutions/495655/wan-quan-er-cha-shu-de-jie-dian-ge-shu-by-leetco-2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
