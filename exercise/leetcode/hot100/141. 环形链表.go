package hot100

/**
leetcode/linked_list/141. 环形链表.go
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
快慢指针，快的每次两步
*/
/**
给你一个链表的头节点 head ，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。

如果链表中存在环 ，则返回 true 。 否则，返回 false 。
*/
func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
	}
	return false
}

/**
执行用时分布
7
ms
击败
48.86%
复杂度分析
消耗内存分布
6.13
MB
击败
17.46%
*/
