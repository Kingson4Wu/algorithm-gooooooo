package hot100

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

/**
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

输入：root = [1,2,3,null,null,4,5]
输出：[1,2,3,null,null,4,5]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]
示例 4：

输入：root = [1,2]
输出：[1,2]

*/

/**
按照之前做题的回忆，树可以通过先序和中序来构造

1、将树序列化 先序和中序字符串，数字之间使用-分割，先序和中序之间使用=分割
2、将字符串反序列化成树，中序的第一个即根节点，可以把先序分成两半，并算出左右子树的长度
*/

type Codec struct {
}

func Constructor2() Codec {

	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

	var firstOrder []string
	/** 先序 */
	var fOrder func(root *TreeNode)
	fOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			fOrder(root.Left)
		}
		firstOrder = append(firstOrder, strconv.Itoa(root.Val))
		if root.Right != nil {
			fOrder(root.Right)
		}
	}
	fOrder(root)

	var middleOrder []string
	/** 中序 */
	var mOrder func(root *TreeNode)
	mOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		middleOrder = append(middleOrder, strconv.Itoa(root.Val))
		if root.Left != nil {
			mOrder(root.Left)
		}
		if root.Right != nil {
			mOrder(root.Right)
		}
	}
	mOrder(root)

	return strings.Join(firstOrder, "-") + "=" + strings.Join(middleOrder, "-")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	var buildTree func(fOrder, mOrder []int) *TreeNode
	buildTree = func(fOrder, mOrder []int) *TreeNode {
		if len(fOrder) == 0 {
			return nil
		}
		root := &TreeNode{
			Val: mOrder[0],
		}
		leftLength := 0
		for _, v := range fOrder {
			if v == mOrder[0] {
				break
			}
			leftLength++
		}

		root.Left = buildTree(fOrder[:leftLength], mOrder[1:1+leftLength])
		root.Right = buildTree(fOrder[leftLength+1:], mOrder[1+leftLength:])
		return root
	}

	arr := strings.Split(data, "=")
	fOrder := strings.Split(arr[0], "-")
	mOrder := strings.Split(arr[1], "-")
	var fO []int
	var mO []int
	for _, v := range fOrder {
		val, _ := strconv.Atoi(v)
		fO = append(fO, val)
	}
	for _, v := range mOrder {
		val, _ := strconv.Atoi(v)
		mO = append(mO, val)
	}

	return buildTree(fO, mO)
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

}
