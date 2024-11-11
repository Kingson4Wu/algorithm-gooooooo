package tree

import "testing"

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	//每一层的开始结点，当pre是nil的时候记录，说明是第一个，后续用于找下一层级
	start := root

	for start != nil {
		//每一层的下一个结点，用于遍历该层的所有结点，利用已经遍历过的next指针找同层级下一个结点
		callStart := start
		start = nil

		var pre *Node

		for callStart != nil {

			if callStart.Left != nil {
				if pre == nil {
					pre = callStart.Left
					start = pre
				} else {
					pre.Next = callStart.Left
					pre = callStart.Left
				}
			}

			if callStart.Right != nil {
				if pre == nil {
					pre = callStart.Right
					start = pre
				} else {
					pre.Next = callStart.Right
					pre = callStart.Right
				}
			}

			callStart = callStart.Next
		}

	}

	return root
}

/**

自己做的！

执行用时：
0 ms
, 在所有 Go 提交中击败了
100.00%
的用户
内存消耗：
6 MB
, 在所有 Go 提交中击败了
100.00%
的用户
通过测试用例：
55 / 55
*/

func TestConnect(t *testing.T) {
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}
	root.Right.Left = &Node{Val: 6}
	root.Right.Right = &Node{Val: 7}

	connect(root)
}
