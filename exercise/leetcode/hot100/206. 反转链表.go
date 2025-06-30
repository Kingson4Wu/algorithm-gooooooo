package hot100

/**
nowcoder/BM1 反转链表.go

时间
4 ms
击败
37.39%
内存
2.4 MB
击败
50.39%
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var pre, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

/**
这次做怎么这么顺利，很快做完

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.25
MB
击败
72.03%

*/
