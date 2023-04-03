package hot100

/**
自己的做法超出时间限制了，重写看了题解
1、使用hash（太粗暴，符合这个题简单级别的定位）
2、双指针，同时走，到空了之后分别指向对方的头指针（这样两边的长度就一样了）

时间
24 ms
击败
94.96%
内存
6.5 MB
击败
98.40%

当链表 headA 和 headB 都不为空时，创建两个指针 pA 和 pB，初始时分别指向两个链表的头节点 headA 和 headB，然后将两个指针依次遍历两个链表的每个节点。具体做法如下：

每步操作需要同时更新指针 pA 和 pB。

如果指针 pA 不为空，则将指针 pA 移到下一个节点；如果指针 pB 不为空，则将指针 pB 移到下一个节点。

如果指针 pA 为空，则将指针 pA 移到链表 headB 的头节点；如果指针 pB 为空，则将指针 pB 移到链表 headA 的头节点。

当指针 pA 和 pB 指向同一个节点或者都为空时，返回它们指向的节点或者 null。

作者：力扣官方题解
链接：https://leetcode.cn/problems/intersection-of-two-linked-lists/solutions/811625/xiang-jiao-lian-biao-by-leetcode-solutio-a8jn/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

/**
简单级别的题目，没会，看了题解，有了启发，但是题解太复杂，我用自己的解法来

1、双指针分别指向A和B，同时移动，先到达的，等待未到达的，计算等待的距离
2、双指针重新开始，慢的先走之前计算的距离，然后同时走，当节点相同时，则为相交的结点

好像不用那么复杂， 第一步优化成分别算出各自的长度即可

超出时间限制 !!!

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}

	pA, pB := headA, headB
	lenA, lenB := 0, 0

	for {
		lenA++
		if pA.Next == nil {
			break
		}
	}

	for {
		lenB++
		if pB.Next == nil {
			break
		}
	}

	if pA != pB {
		return nil
	}

	pA, pB = headA, headB
	if lenA > lenB {
		preGo := lenA - lenB
		for i := 0; i < preGo; i++ {
			pA = pA.Next
		}
	} else {
		preGo := lenB - lenA
		for i := 0; i < preGo; i++ {
			pB = pB.Next
		}
	}

	for {
		if pA == pB {
			return pA
		}
		pA = pA.Next
		pB = pB.Next
	}

	return nil
}*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}

	pA, pB := headA, headB

	for {
		if pA == nil && pB == nil {
			return nil
		}
		if pA == pB {
			return pA
		}

		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return nil
}
