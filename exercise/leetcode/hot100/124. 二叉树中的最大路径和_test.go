package hot100

import (
	"fmt"
	"testing"
)

/**
递归写完了。。。
没考虑根结点是负数的情况。。

改了三次才写对

时间
164 ms
击败
6.81%
内存
6.9 MB
击败
86.67%

答案也是递归，比较简洁

func maxPathSum(root *TreeNode) int {
    maxSum := math.MinInt32
    var maxGain func(*TreeNode) int
    maxGain = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        // 递归计算左右子节点的最大贡献值
        // 只有在最大贡献值大于 0 时，才会选取对应子节点
        leftGain := max(maxGain(node.Left), 0)
        rightGain := max(maxGain(node.Right), 0)

        // 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
        priceNewPath := node.Val + leftGain + rightGain

        // 更新答案
        maxSum = max(maxSum, priceNewPath)

        // 返回节点的最大贡献值
        return node.Val + max(leftGain, rightGain)
    }
    maxGain(root)
    return maxSum
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/binary-tree-maximum-path-sum/solutions/297005/er-cha-shu-zhong-de-zui-da-lu-jing-he-by-leetcode-/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

/**
输入
root =
[-3]
添加到测试用例
输出
0
预期结果
-3

输入
root =
[1,-2,-3,1,3,-2,null,-1]
添加到测试用例
输出
4
预期结果
3

root =
[5,4,8,11,null,13,4,7,2,null,null,null,1]
添加到测试用例
输出
55
预期结果
48
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {

	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	maxVal := root.Val
	if root.Left == nil {
		right := maxPathSum2(root.Right)

		if right > maxVal {
			maxVal = right
		}
		if right+root.Val > maxVal {
			maxVal = right + root.Val
		}

		right2 := maxPathSum(root.Right)
		if right2 > maxVal {
			maxVal = right2
		}
		return maxVal
	}

	if root.Right == nil {
		left := maxPathSum2(root.Left)

		if left > maxVal {
			maxVal = left
		}
		if left+root.Val > maxVal {
			maxVal = left + root.Val
		}

		left2 := maxPathSum(root.Left)
		if left2 > maxVal {
			maxVal = left2
		}
		return maxVal
	}

	right := maxPathSum2(root.Right)

	if right > maxVal {
		maxVal = right
	}
	if right+root.Val > maxVal {
		maxVal = right + root.Val
	}

	right2 := maxPathSum(root.Right)
	if right2 > maxVal {
		maxVal = right2
	}

	left := maxPathSum2(root.Left)

	if left > maxVal {
		maxVal = left
	}
	if left+root.Val > maxVal {
		maxVal = left + root.Val
	}

	left2 := maxPathSum(root.Left)
	if left2 > maxVal {
		maxVal = left2
	}

	if left+right+root.Val > maxVal {
		maxVal = left + right + root.Val
	}

	return maxVal
}

/** 包含root结点，但是只包含左或右的路径 ！！！ */
func maxPathSum2(root *TreeNode) int {

	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	maxVal := root.Val
	if root.Left == nil {
		right := maxPathSum2(root.Right)
		if right+root.Val > maxVal {
			maxVal = right + root.Val
		}
		return maxVal
	}
	if root.Right == nil {
		left := maxPathSum2(root.Left)
		if left+root.Val > maxVal {
			maxVal = left + root.Val
		}
		return maxVal
	}

	left := maxPathSum2(root.Left)
	right := maxPathSum2(root.Right)

	if left+root.Val > maxVal {
		maxVal = left + root.Val
	}
	if right+root.Val > maxVal {
		maxVal = right + root.Val
	}

	return maxVal

}

/** 这个写法是错的，没考虑要相连 */
/*func maxPathSum(root *TreeNode) int {

	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	if root.Left == nil {
		right := maxPathSum(root.Right)

		maxVal := root.Val

		if right > maxVal {
			maxVal = right
		}

		if right+root.Val > maxVal {
			maxVal = right + root.Val
		}
		return maxVal
	}
	if root.Right == nil {
		left := maxPathSum(root.Left)

		maxVal := root.Val

		if left > maxVal {
			maxVal = left
		}

		if left+root.Val > maxVal {
			maxVal = left + root.Val
		}
		return maxVal
	}

	left := maxPathSum(root.Left)
	right := maxPathSum(root.Right)

	maxVal := root.Val

	if left > maxVal {
		maxVal = left
	}
	if right > maxVal {
		maxVal = right
	}

	if left+root.Val > maxVal {
		maxVal = left + root.Val
	}
	if right+root.Val > maxVal {
		maxVal = right + root.Val
	}
	if left+right+root.Val > maxVal {
		maxVal = left + right + root.Val
	}

	return maxVal
}*/

func TestMaxPathSum(t *testing.T) {

	fmt.Println(maxPathSum(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}))

	//[1,-2,-3,1,3,-2,null,-1]
	fmt.Println(maxPathSum(&TreeNode{Val: 1, Left: &TreeNode{Val: -2, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: -1}}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: -3, Left: &TreeNode{Val: -2}}}))

}
