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

提示：

树中结点数在范围 [0, 104] 内
-1000 <= Node.val <= 1000

*/

/**
按照之前做题的回忆，树可以通过先序和中序来构造

1、将树序列化 先序和中序字符串，数字之间使用-分割，先序和中序之间使用=分割
2、将字符串反序列化成树，中序的第一个即根节点，可以把先序分成两半，并算出左右子树的长度

没有说 数值都不相等？？

[4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2]

测试用例就存在相同的数值！！
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

		/**
		判断是否存在左或右， 边界条件老是忘！！！
		*/
		if leftLength > 0 {
			root.Left = buildTree(fOrder[:leftLength], mOrder[1:leftLength+1])
		}
		if len(fOrder)-1-leftLength > 0 {
			root.Right = buildTree(fOrder[leftLength+1:], mOrder[1+leftLength:])
		}
		return root
	}

	arr := strings.Split(data, "=")

	var fOrder []string
	if len(arr[0]) > 0 {
		fOrder = strings.Split(arr[0], "-")
		// 空白字符串Split会出现长度位1的数组，值为"" !!!!!!!
	}
	var mOrder []string
	if len(arr[1]) > 0 {
		mOrder = strings.Split(arr[1], "-")
	}

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

	data = ser.serialize(nil)
	fmt.Println(data)
	ans = deser.deserialize(data)
	fmt.Println(ans)

	s := "a"
	fmt.Println(s[1:1])

}

/**

panic: runtime error: slice bounds out of range [2:1]
main.(*Codec).deserialize.func1({0xc00008c028, 0x1, 0x1b}, {0xc00008c150, 0x1, 0x16})
solution.go, line 78
main.(*Codec).deserialize.func1({0xc00008c028, 0x1, 0x1b}, {0xc00008c148, 0x1, 0x17})
solution.go, line 77
main.(*Codec).deserialize.func1({0xc00008c028, 0x5, 0x1b}, {0xc00008c140, 0x5, 0x18})
solution.go, line 77
main.(*Codec).deserialize.func1({0xc00008c028, 0x5, 0x1b}, {0xc00008c138, 0x5, 0x19})
solution.go, line 77
main.(*Codec).deserialize.func1({0xc00008c028, 0x13, 0x1b}, {0xc00008c130, 0x13, 0x1a})
solution.go, line 77
main.(*Codec).deserialize.func1({0xc00008c020, 0x14, 0x1c}, {0xc00008c128, 0x14, 0x1b})
solution.go, line 78
main.(*Codec).deserialize.func1({0xc00008c020, 0x19, 0x1c}, {0xc00008c120, 0x19, 0x1c})
solution.go, line 77
main.(*Codec).deserialize.func1({0xc00008c018, 0x1a, 0x1d}, {0xc00008c118, 0x1a, 0x1d})
solution.go, line 78
main.(*Codec).deserialize.func1({0xc00008c000, 0x1d, 0x20}, {0xc00008c100, 0x1d, 0x20})
solution.go, line 78
main.(*Codec).deserialize(0x4b2be0?, {0xc000056060?, 0xc00000ca38?})
solution.go, line 104
main.__helper__(0xc00003e070?)
solution.go, line 121
main.main()
solution.go, line 148
最后执行的输入
添加到测试用例
[4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2]
*/

/**
panic: runtime error: slice bounds out of range [:2] with capacity 1
main.(*Codec).deserialize.func1({0xc00008a088, 0x1, 0xf}, {0xc00008a1f8, 0x1, 0x1})
*/
