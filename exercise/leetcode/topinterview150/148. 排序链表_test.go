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
自己做，只能用冒泡排序了
*/
/**
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
*/
/*func sortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var pre *ListNode
	newHead := head
	var tail *ListNode
	for head != nil && head.Next != nil {
		i := head
		j := head.Next
		pre = nil
		for i != nil && j != nil && j != tail {
			if i.Val > j.Val {
				if pre != nil {
					pre.Next = j
				} else {
					newHead = j
				}
				i.Next = j.Next
				j.Next = i
				pre = j

				j = i.Next
				continue
			}
			i, j = i.Next, j.Next
		}
		if tail == nil {
			tail = i
		}
		head = newHead
	}

	return newHead
}*/

/*func sortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	return head
}*/

/**
真恶心，答案都看吐了，各种指针交换
*/

func sortList(head *ListNode) *ListNode {
	length := 0
	origin := head
	for head != nil {
		length++
		head = head.Next
	}
	head = origin
	for i := 1; i < length; i <<= 1 {
		// i 是子链表长度

		var newHead *ListNode
		var tail *ListNode
		for k := 0; length-k*i*2 > i; k++ {
			h1 := head
			h2 := head
			for j := 0; j < i; j++ {
				h2 = h2.Next
			}
			// h3 下一个开始的点
			h3 := h2

			for j := 0; j < i && h3 != nil; j++ {
				h3 = h3.Next
			}
			h, t := merge(h1, h2, i)
			if newHead == nil {
				newHead = h
			} else {
				tail.Next = h
			}
			tail = t
			head = h3
		}
		/*if i*2 >= length/2 && head != nil {
			newHead, _ = merge(newHead, head, i*2)
		}*/
		/*if i<<2 > length/2 {
			if head != nil {
				newHead = merge(newHead, head, 0)
			}
		}*/
		head = newHead
	}

	return head
}

func merge(head1 *ListNode, head2 *ListNode, length int) (*ListNode, *ListNode) {

	var head, pre *ListNode
	head1Length, head2Length := length, length
	for head1 != nil && head2 != nil && head1Length > 0 && head2Length > 0 {
		if head1.Val <= head2.Val {
			if head == nil {
				head = head1
				pre = head1
			} else {
				pre.Next = head1
				pre = pre.Next
			}
			head1 = head1.Next
			head1Length--
		} else {
			if head == nil {
				head = head2
				pre = head2
			} else {
				pre.Next = head2
				pre = pre.Next
			}
			head2 = head2.Next
			head2Length--
		}
	}
	for head1 != nil && head1Length > 0 {
		pre.Next = head1
		pre = pre.Next
		head1 = head1.Next
		head1Length--
		/*if head1Length == 0 {
			head1.Next = nil
		}*/
	}
	for head2 != nil && head2Length > 0 {
		pre.Next = head2
		pre = pre.Next
		head2 = head2.Next
		head2Length--
	}
	// 结束前把整个链表重新连起来
	tail := pre
	pre.Next = head2
	return head, tail

}

func TestSortList(t *testing.T) {
	head := sortList(&ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}})
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	fmt.Println("====")
	head = sortList(&ListNode{Val: -1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 0}}}}})
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

	fmt.Println("====")
	head = sortList(nil)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	fmt.Println("====")
	head = sortList(&ListNode{Val: -1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 0, Next: &ListNode{Val: -3, Next: &ListNode{Val: 8}}}}}}})
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

}

/**
总体思想，分治和归并排序
定义merge函数，用于两两合并已排序的链表

*/
/**

写得不好，链表要各种debug指针变量

执行用时分布
27
ms
击败
21.64%
复杂度分析
消耗内存分布
8.77
MB
击败
94.79%

*/
