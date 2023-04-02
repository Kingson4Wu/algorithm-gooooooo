package hot100

import (
	"fmt"
	"testing"
)

/**

个人目前做法：

快慢指针找中点
同时慢指针构造逆转列表  (逆转列表可以原地反转，不用额外的数组空间)
从中点开始和逆转列表对比

时间
120 ms
击败
70.79%
内存
8.5 MB
击败
74.32%
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head.Next

	var prefix []int
	prefix = append(prefix, slow.Val)

	for fast.Next != nil {

		if fast.Next.Next == nil {
			slow = slow.Next
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
		prefix = append(prefix, slow.Val)
	}

	slow = slow.Next
	i := len(prefix) - 1
	for slow != nil {
		if slow.Val != prefix[i] {
			return false
		}
		i--
		slow = slow.Next
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	fmt.Println(isPalindrome(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}))
	fmt.Println(isPalindrome(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}}))
	fmt.Println(isPalindrome(&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}}))
}
