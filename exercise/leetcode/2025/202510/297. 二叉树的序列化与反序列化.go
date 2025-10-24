package _02510

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {

	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var arr []string
	var build func(root *TreeNode)
	build = func(root *TreeNode) {
		if root == nil {
			arr = append(arr, "null")
			return
		}
		arr = append(arr, strconv.Itoa(root.Val))
		build(root.Left)
		build(root.Right)
	}
	build(root)
	return strings.Join(arr, ",")

}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	arr := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if arr[0] == "null" {
			arr = arr[1:]
			return nil
		}
		v, _ := strconv.Atoi(arr[0])
		arr = arr[1:]
		root := &TreeNode{Val: v}
		root.Left = build()
		root.Right = build()
		return root
	}
	return build()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

/**
根据回忆写对了，关键是递归，按顺序遍历或构造，递归函数外有全局变量构造或消耗

执行用时分布
7
ms
击败
75.52%
复杂度分析
消耗内存分布
8.77
MB
击败
45.83%
*/
