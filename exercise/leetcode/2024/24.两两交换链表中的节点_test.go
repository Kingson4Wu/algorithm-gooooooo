package _024

import (
	"fmt"
	"testing"
)

// 自己做的，搞清楚交换的关系，定义对应的各种指针

func swapPairs(head *ListNode) *ListNode {

	var newHead, preTail, nextHead, first, second *ListNode

	for head != nil {
		first = head
		second = head.Next
		if second == nil {
			break
		}
		nextHead = second.Next

		first.Next = nextHead
		second.Next = first
		if newHead == nil {
			newHead = second
		}
		if preTail != nil {
			preTail.Next = second
		}
		preTail = first
		head = nextHead
	}
	if newHead == nil {
		newHead = head
	}

	return newHead
}

func TestSwapPairs(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	head = swapPairs(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	fmt.Println("=============")
	head = nil
	head = swapPairs(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	fmt.Println("=============")
	head = &ListNode{Val: 1}
	head = swapPairs(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	fmt.Println("=============")
	head = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	head = swapPairs(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

}
