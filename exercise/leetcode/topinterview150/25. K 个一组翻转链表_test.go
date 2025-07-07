package topinterview150

import (
	"fmt"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

nowcoder/BM3 链表中的节点每k个一组翻转.go
*/
/**
自己写得，最终还是要debug调整
定义一个逆序的方法，然后每k个传进去，注意tail置为nil
*/
func reverseKGroup(head *ListNode, k int) *ListNode {

	var newHead, pre, tail, next *ListNode
	count := 0
	pre = head
	for head != nil {
		count++
		if count == k {
			next = head.Next
			head.Next = nil
			h, t := reverse(pre)
			if newHead == nil {
				newHead = h
			} else {
				tail.Next = h
			}
			pre = next
			tail = t
			count = 0
			head = next
		} else {
			head = head.Next
		}
	}
	if count > 0 {
		tail.Next = pre
	}
	return newHead
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	var pre, next, tail *ListNode
	for head != nil {
		if tail == nil {
			tail = head
		}
		next = head.Next
		if pre != nil {
			head.Next = pre
		}
		pre = head
		head = next
	}
	if tail != nil {
		tail.Next = nil
	}
	return pre, tail
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
5.17
MB
击败
99.91%
*/

func TestReverseKGroup(t *testing.T) {
	h := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	fmt.Println(reverseKGroup(h, 2))
}
