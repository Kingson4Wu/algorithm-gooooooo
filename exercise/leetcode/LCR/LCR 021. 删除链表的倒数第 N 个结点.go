package LCR

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
*
执行用时分布
2
ms
击败
49.47%
复杂度分析
消耗内存分布
3.90
MB
击败
53.68%
复杂度分析

根据回忆完成
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur, next := head, head
	for i := 0; i < n; i++ {
		next = next.Next
	}
	var pre *ListNode
	for next != nil {
		pre = cur
		cur = cur.Next
		next = next.Next
	}
	if pre == nil {
		return cur.Next
	}
	pre.Next = cur.Next
	return head
}
