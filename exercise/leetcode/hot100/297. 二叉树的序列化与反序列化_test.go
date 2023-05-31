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

没有说 数值都不相等！！！ 不能用先序和中序的方法
按照之前做题的回忆，树可以通过先序和中序来构造

使用先序和中序的方法的前提是：
1、数值不相同
2、数字是连续的，没有null的字符串可以表示空

这道题可以使用先序来序列化和反序列化！！！
1、使用null表示结点为空
2、使用递归
3、使用全局变量！！！简化递归的传参问题！！要记住递归运行时也是串行的！！


测试用例就存在相同的数值！！

时间
8 ms
击败
87.39%
内存
6.6 MB
击败
53.96%
*/

type Codec struct {
}

func Constructor2() Codec {

	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

	var result []string

	var serialize func(*TreeNode)
	serialize = func(root *TreeNode) {
		if root == nil {
			result = append(result, "null")
			return
		}
		result = append(result, strconv.Itoa(root.Val))
		serialize(root.Left)
		serialize(root.Right)
	}

	serialize(root)

	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	if data == "null" {
		return nil
	}
	result := strings.Split(data, ",")

	var deserialize func() *TreeNode
	deserialize = func() *TreeNode {
		if len(result) == 0 {
			return nil
		}
		if result[0] == "null" {
			result = result[1:]
			return nil
		}
		v, _ := strconv.Atoi(result[0])
		root := &TreeNode{
			Val: v,
		}
		result = result[1:]
		root.Left = deserialize()
		root.Right = deserialize()

		return root
	}
	root := deserialize()
	return root
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

[4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2]
*/
