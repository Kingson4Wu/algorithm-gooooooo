package tencent50

import "testing"

/**
题解的空间复杂度是O(1)

方法一：闭合为环
思路及算法

记给定链表的长度为 nnn，注意到当向右移动的次数 k≥n 时，我们仅需要向右移动 k mod n 次即可。因为每 n 次移动都会让链表变为原状。这样我们可以知道，新链表的最后一个节点为原链表的第 (n - 1) - (k mod n) 个节点（从 0 开始计数）。

这样，我们可以先将给定的链表连接成环，然后将指定位置断开。

具体代码中，我们首先计算出链表的长度 n，并找到该链表的末尾节点，将其与头节点相连。这样就得到了闭合为环的链表。然后我们找到新链表的最后一个节点（即原链表的第  (n - 1) - (k mod n)  个节点），将当前闭合为环的链表断开，即可得到我们所需要的结果。

特别地，当链表长度不大于 1，或者 k 为 n 的倍数时，新链表将与原链表相同，我们无需进行任何处理。

作者：力扣官方题解
链接：https://leetcode.cn/problems/rotate-list/solutions/681812/xuan-zhuan-lian-biao-by-leetcode-solutio-woq1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

/**
注意是右移不是左移

个人想法：
1、遍历一趟，得到逆序列表（hash保存）和长度length
2、算出k%length
3、从尾部遍历，找到变更的位置

时间
0 ms
击败
100%
内存
2.8 MB
击败
5.14%
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func rotateRight(head *ListNode, k int) *ListNode {

	if k == 0 || head == nil {
		return head
	}

	m := map[*ListNode]*ListNode{}

	length := 0
	p := head
	var tail *ListNode
	for p != nil {
		tail = p
		length++
		m[p.Next] = p
		p = p.Next
	}

	k %= length

	if k == 0 {
		return head
	}

	var pre *ListNode
	cur := tail
	for i := 0; i < k; i++ {
		pre = cur
		cur = m[cur]
	}
	cur.Next = nil

	p = pre
	for p.Next != nil {
		p = p.Next
	}
	p.Next = head

	return pre
}

/*func rotateRight(head *ListNode, k int) *ListNode {

	if k == 0 || head == nil {
		return head
	}

	var pre *ListNode
	cur := head

	length := 0

	for i := 0; i < k; i++ {
		if cur.Next == nil {
			break
		}
		pre = cur
		cur = cur.Next
		length++
	}

	//说明长度为1 *
	if length == 0 {
		return head
	}

	//if length == k {
		//没走到链表尽头

	//}
	if length == k {
		// 长度为length+1，已经走了length步，还有k-length步没走
		k %= length + 1
		for i := 0; i < k; i++ {
			pre = cur
			cur = cur.Next
		}
	}

	pre.Next = nil

	p := cur
	for p.Next != nil {
		p = p.Next
	}
	p.Next = head

	return cur
}*/

func TestRotateRight(t *testing.T) {
	rotateRight(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 2)
}

/**
输入
head =
[1,2,3,4,5]
k =
2
输出
[5,1,2,3,4]
预期结果
[4,5,1,2,3]
*/
