package hot100

import (
	"fmt"
	"testing"
)

/**

看了题解才会做

注意！快慢指针要设置同一点出发

fast = a + n(b+c) +b
slow = a+b

a + n (b+c) +b = 2 (a+b)

a = c+ (n-1)(b+c)

时间
0 ms
击败
100%
内存
3.5 MB
击败
82.19%

*/

/**
给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

不允许修改 链表。
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	slow, fast := head, head

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next
		if fast == nil {
			return nil
		}
		if slow == fast {
			break
		}
	}

	fast = head
	for fast != nil {
		if slow == fast {
			break
		}
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}

func TestDetectCycle(t *testing.T) {
	roo1 := &ListNode{Val: 1}
	roo2 := &ListNode{Val: 2, Next: roo1}
	roo1.Next = roo2
	fmt.Println(detectCycle(roo1).Val)

	roo3 := &ListNode{Val: 1}
	roo3.Next = roo3
	fmt.Println(detectCycle(roo3).Val)
}
