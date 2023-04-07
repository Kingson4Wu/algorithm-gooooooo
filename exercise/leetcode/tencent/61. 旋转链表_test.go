package tencent

import "testing"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func rotateRight(head *ListNode, k int) *ListNode {

	return nil
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
