package topinterview150

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
nowcoder/BM15 删除有序链表中重复的元素-I.go
*/
func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	pre := head
	head = head.Next
	for head != nil {
		if head.Val == pre.Val {
			
		}
	}

	return head
}
