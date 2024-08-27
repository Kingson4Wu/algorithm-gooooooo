package _024

import (
	"fmt"
	"testing"
)

// 自己做的，搞清楚交换的关系，定义对应的各种指针

// 官方用递归简单多了 ！但递归耗时也会较长
// 官方的迭代用的变量比较少，但个人感觉没有我写的比较好理解，虽然比较啰嗦，但变量都定义得比较明确

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
2.13
MB
击败
58.59%

*/

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

/**
官方方法用递归，看起来简短
可以通过递归的方式实现两两交换链表中的节点。
递归的终止条件是链表中没有节点，或者链表中只有一个节点，此时无法进行交换

func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newHead := head.Next
    head.Next = swapPairs(newHead.Next)
    newHead.Next = head
    return newHead
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/swap-nodes-in-pairs/solutions/444474/liang-liang-jiao-huan-lian-biao-zhong-de-jie-di-91/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

/**
迭代
func swapPairs(head *ListNode) *ListNode {
    dummyHead := &ListNode{0, head}
    temp := dummyHead
    for temp.Next != nil && temp.Next.Next != nil {
        node1 := temp.Next
        node2 := temp.Next.Next
        temp.Next = node2
        node1.Next = node2.Next
        node2.Next = node1
        temp = node1
    }
    return dummyHead.Next
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/swap-nodes-in-pairs/solutions/444474/liang-liang-jiao-huan-lian-biao-zhong-de-jie-di-91/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
