package _025

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
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

func Constructor2() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var nodes []string
	if root == nil {
		return "null"
	}
	nodes = append(nodes, strconv.Itoa(root.Val))
	left := this.serialize(root.Left)
	right := this.serialize(root.Right)
	nodes = append(nodes, left)
	nodes = append(nodes, right)
	return strings.Join(nodes, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	if len(data) == 0 {
		return nil
	}
	nodes := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if len(nodes) == 0 || nodes[0] == "null" {
			nodes = nodes[1:]
			return nil
		}
		v, _ := strconv.Atoi(nodes[0])
		root := &TreeNode{Val: v}
		nodes = nodes[1:]
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

func TestSerialize(t *testing.T) {

	ser := Constructor2()
	deser := Constructor2()
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 5,
		},
		Right: &TreeNode{
			Val:   9,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
		},
	}
	data := ser.serialize(root)
	fmt.Println(data)
	ans := deser.deserialize(data)
	fmt.Println(ans)

	data = ser.serialize(nil)
	fmt.Println(data)
	ans = deser.deserialize(data)
	fmt.Println(ans)

	s := "a"
	fmt.Println(s[1:1])

}

/**
4,5,null,null,9,3,null,null,1,null,2,null,null
&{4 0x140000ac168 0x140000ac180}
null
<nil>
*/

/**
自己写得，差的没写出来
序列化和反序列化都是递归
反序列化定义一个内部方法来递归，再定一个全局变量，在递归中逐渐消耗这个全局变量， 不用在递归里传这么麻烦

执行用时分布
15
ms
击败
27.27%
复杂度分析
消耗内存分布
9.79
MB
击败
15.15%
复杂度分析

*/
