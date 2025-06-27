package GNode

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */
/**
给你一个 n * n 矩阵 grid ，矩阵由若干 0 和 1 组成。请你用四叉树表示该矩阵 grid 。

你需要返回能表示矩阵 grid 的 四叉树 的根结点。

四叉树数据结构中，每个内部节点只有四个子节点。此外，每个节点都有两个属性：

val：储存叶子结点所代表的区域的值。1 对应 True，0 对应 False。注意，当 isLeaf 为 False 时，你可以把 True 或者 False 赋值给节点，两种值都会被判题机制 接受 。
isLeaf: 当这个节点是一个叶子结点时为 True，如果它有 4 个子节点则为 False 。
class Node {
    public boolean val;
    public boolean isLeaf;
    public Node topLeft;
    public Node topRight;
    public Node bottomLeft;
    public Node bottomRight;
}
我们可以按以下步骤为二维区域构建四叉树：

如果当前网格的值相同（即，全为 0 或者全为 1），将 isLeaf 设为 True ，将 val 设为网格相应的值，并将四个子节点都设为 Null 然后停止。
如果当前网格的值不同，将 isLeaf 设为 False， 将 val 设为任意值，然后如下图所示，将当前网格划分为四个子网格。
使用适当的子网格递归每个子节点。


如果你想了解更多关于四叉树的内容，可以参考 wiki 。

四叉树格式：

你不需要阅读本节来解决这个问题。只有当你想了解输出格式时才会这样做。输出为使用层序遍历后四叉树的序列化形式，其中 null 表示路径终止符，其下面不存在节点。

它与二叉树的序列化非常相似。唯一的区别是节点以列表形式表示 [isLeaf, val] 。

如果 isLeaf 或者 val 的值为 True ，则表示它在列表 [isLeaf, val] 中的值为 1 ；如果 isLeaf 或者 val 的值为 False ，则表示值为 0 。


*/

/**

自己思考写的， 分治递归
*/

func simplifyNode(node *Node) {
	if node.TopLeft != nil && node.TopLeft.IsLeaf && node.TopRight != nil && node.TopRight.IsLeaf && node.BottomLeft != nil && node.BottomLeft.IsLeaf && node.BottomRight != nil && node.BottomRight.IsLeaf {
		if node.TopLeft.Val && node.TopRight.Val && node.BottomLeft.Val && node.BottomRight.Val {
			node.Val = true
			node.IsLeaf = true
			node.TopLeft = nil
			node.TopRight = nil
			node.BottomLeft = nil
			node.BottomRight = nil
		} else if !node.TopLeft.Val && !node.TopRight.Val && !node.BottomLeft.Val && !node.BottomRight.Val {
			node.Val = false
			node.IsLeaf = true
			node.TopLeft = nil
			node.TopRight = nil
			node.BottomLeft = nil
			node.BottomRight = nil
		}
	}
}

func construct(grid [][]int) *Node {

	var createNode func(grid [][]int, centreX, centreY, size int) *Node
	createNode = func(grid [][]int, centreX, centreY, size int) *Node {

		if size == 1 {
			if grid[centreX][centreY] == 1 {
				return &Node{Val: true, IsLeaf: true}
			} else {
				return &Node{Val: false, IsLeaf: true}
			}
		}

		if size == 2 {
			node := &Node{Val: true}
			if grid[centreX][centreY] == 1 {
				node.TopLeft = &Node{Val: true, IsLeaf: true}
			} else {
				node.TopLeft = &Node{Val: false, IsLeaf: true}
			}
			if grid[centreX][centreY+1] == 1 {
				node.TopRight = &Node{Val: true, IsLeaf: true}
			} else {
				node.TopRight = &Node{Val: false, IsLeaf: true}
			}
			if grid[centreX+1][centreY] == 1 {
				node.BottomLeft = &Node{Val: true, IsLeaf: true}
			} else {
				node.BottomLeft = &Node{Val: false, IsLeaf: true}
			}
			if grid[centreX+1][centreY+1] == 1 {
				node.BottomRight = &Node{Val: true, IsLeaf: true}
			} else {
				node.BottomRight = &Node{Val: false, IsLeaf: true}
			}
			simplifyNode(node)
			return node
		}
		node := &Node{Val: true}
		node.TopLeft = createNode(grid, centreX, centreY, size/2)
		node.TopRight = createNode(grid, centreX, centreY+size/2, size/2)
		node.BottomLeft = createNode(grid, centreX+size/2, centreY, size/2)
		node.BottomRight = createNode(grid, centreX+size/2, centreY+size/2, size/2)
		simplifyNode(node)
		return node
	}

	return createNode(grid, 0, 0, len(grid))
}

/**
执行用时分布
3
ms
击败
93.56%
复杂度分析
消耗内存分布
8.01
MB
击败
34.16%
*/
