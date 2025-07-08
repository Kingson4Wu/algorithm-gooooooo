package topinterview145

import "testing"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
*
给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别分组，保持它们原有的相对顺序，然后把偶数索引节点分组连接到奇数索引节点分组之后，返回重新排序的链表。

第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。

请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。

你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题。

示例 1:

输入: head = [1,2,3,4,5]
输出: [1,3,5,2,4]
示例 2:

输入: head = [2,1,3,5,6,4,7]
输出: [2,3,6,7,1,5,4]

提示:

n ==  链表中的节点数
0 <= n <= 104
-106 <= Node.val <= 106
*/
func oddEvenList(head *ListNode) *ListNode {
	var evenHead, evenNext, oddHead, oddNext *ListNode
	count := 1
	for head != nil {
		if count%2 == 1 {
			if oddHead == nil {
				oddHead = head
				oddNext = head
			} else {
				oddNext.Next = head
				oddNext = oddNext.Next
			}
		} else {
			if evenHead == nil {
				evenHead = head
				evenNext = head
			} else {
				evenNext.Next = head
				evenNext = evenNext.Next
			}
		}
		count++
		head = head.Next
	}
	if evenHead != nil {
		//这句置空是重点！！！
		evenNext.Next = nil
		oddNext.Next = evenHead
	}
	return oddHead
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.86
MB
击败
76.14%
*/

func TestOddEvenList(t *testing.T) {
	oddEvenList(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}})
}
