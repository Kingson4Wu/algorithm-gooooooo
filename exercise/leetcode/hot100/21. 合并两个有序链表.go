package hot100

/**
自己做的
注意要想清楚，定义必要的指针变量

时间
4 ms
击败
46.57%
内存
2.4 MB
击败
72.69%
*/
/**
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	root1, root2 := list1, list2
	var root, rootNext *ListNode
	if list1.Val <= list2.Val {
		root = list1
		root1 = root1.Next
	} else {
		root = list2
		root2 = root2.Next
	}
	rootNext = root

	for root1 != nil || root2 != nil {

		if root1 == nil {
			rootNext.Next = root2
			break
		}
		if root2 == nil {
			rootNext.Next = root1
			break
		}

		if root1.Val <= root2.Val {
			rootNext.Next = root1
			rootNext = rootNext.Next
			root1 = root1.Next
			continue
		} else {
			rootNext.Next = root2
			rootNext = rootNext.Next
			root2 = root2.Next
		}
	}
	return root
}
