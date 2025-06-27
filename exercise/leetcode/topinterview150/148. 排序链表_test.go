package topinterview150

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

func sortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	return head
}

/**
真恶心，答案都看吐了，各种指针交换
*/
