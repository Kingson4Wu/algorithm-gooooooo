package hot100

import (
	"fmt"
	"testing"
)

/**
leetcode/linked_list/19. 删除链表的倒数第 N 个结点.go
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head.Next == nil {
		return nil
	}
	var arr []*ListNode
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	index := len(arr) - n
	if index == 0 {
		return arr[1]
	}
	if index == len(arr)-1 {
		arr[len(arr)-2].Next = nil
		return arr[0]
	}
	arr[index-1].Next = arr[index+1]
	return arr[0]
}

func TestRemoveNthFromEnd(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	head = removeNthFromEnd(head, 2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

	fmt.Println("====")
	head = &ListNode{Val: 1}
	head = removeNthFromEnd(head, 1)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

	fmt.Println("====")
	head = &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	head = removeNthFromEnd(head, 1)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

/**
方法三：双指针
思路与算法

我们也可以在不预处理出链表的长度，以及使用常数空间的前提下解决本题。

由于我们需要找到倒数第 n 个节点，因此我们可以使用两个指针 first 和 second 同时对链表进行遍历，并且 first 比 second 超前 n 个节点。当 first 遍历到链表的末尾时，second 就恰好处于倒数第 n 个节点。

作者：力扣官方题解
链接：https://leetcode.cn/problems/remove-nth-node-from-end-of-list/solutions/450350/shan-chu-lian-biao-de-dao-shu-di-nge-jie-dian-b-61/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
