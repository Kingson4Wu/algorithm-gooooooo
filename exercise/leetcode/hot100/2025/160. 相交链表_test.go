package _025

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
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
*
用change变量，写得真丑
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	change := 0

	for {
		if pA == pB {
			return pA
		}
		pA = pA.Next
		pB = pB.Next
		if pA == nil {
			if change == 2 {
				return nil
			}
			pA = headB
			change++
		}
		if pB == nil {
			if change == 2 {
				return nil
			}
			pB = headA
			change++
		}
	}
	return nil
}

/**
执行用时分布
26
ms
击败
25.43%
复杂度分析
消耗内存分布
8.40
MB
击败
71.35%
复杂度分析

*/

func TestGetIntersectionNode(t *testing.T) {
	/*node := &ListNode{Val: 4, Next: &ListNode{Val: 7}}
	headA := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: node}}
	headB := &ListNode{Val: 9, Next: node}
	fmt.Println(getIntersectionNode(headA, headB).Val)*/
	headA := &ListNode{Val: 1, Next: &ListNode{Val: 3}}
	headB := &ListNode{Val: 9}
	fmt.Println(getIntersectionNode(headA, headB).Val)
}
