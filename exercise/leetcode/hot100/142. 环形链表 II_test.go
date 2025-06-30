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

快慢指针相遇之后，快指针从head开始重新走，slow继续走，再次相遇的点即入环点
（方法1使用的是hash，以指针为key）

有了 a=c+(n−1)(b+c) 的等量关系，我们会发现：从相遇点到入环点的距离加上 n−1 圈的环长，恰好等于从链表头部到入环点的距离。

因此，当发现 slow 与 fast 相遇时，我们再额外使用一个指针 ptr。起始，它指向链表头部；随后，它和 slow 每次向后移动一个位置。最终，它们会在入环点相遇。


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

/**
官方的答案比较简洁
func detectCycle(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil {
        slow = slow.Next
        if fast.Next == nil {
            return nil
        }
        fast = fast.Next.Next
        if fast == slow {
            p := head
            for p != slow {
                p = p.Next
                slow = slow.Next
            }
            return p
        }
    }
    return nil
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/linked-list-cycle-ii/solutions/441131/huan-xing-lian-biao-ii-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
