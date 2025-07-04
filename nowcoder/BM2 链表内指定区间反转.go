package nowcoder

/**
 *
 * @param head ListNode类
 * @param m int整型
 * @param n int整型
 * @return ListNode类
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// write code here

	if m == n {
		return head
	}

	index := 0
	indexHead := head

	var previous *ListNode = nil
	var startPrevious *ListNode = nil
	var start *ListNode = nil
	reverse := false

	for indexHead != nil {
		index++

		if !reverse {
			if index == m {
				start = indexHead
				startPrevious = previous
				reverse = true
			}
			previous = indexHead
			indexHead = indexHead.Next
		} else {
			temp := indexHead.Next
			indexHead.Next = previous
			previous = indexHead
			indexHead = temp

			if index == n {

				if startPrevious != nil {
					startPrevious.Next = previous
				} else {
					head = previous
				}

				start.Next = indexHead
				break
			}
		}

	}

	return head
}

/**
将一个节点数为 size 链表 m 位置到 n 位置之间的区间反转，要求时间复杂度 O(n)，空间复杂度 O(1)。
*/

/**
运行时间：4ms
超过54.04% 用Go提交的代码
占用内存：1040KB
超过68.70%用Go提交的代码

/**
自己完成的，不过写得比较久
链表的题目比较简单，但是要提前画好图，理清楚需要哪些变量，这样写起来思路会清晰很多！！！
*/

/**
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n || head == nil {
		return head
	}

	//哨兵节点， 避免处理头指针特殊情况
	dummy := &ListNode{Next: head}
	pre := dummy

	// 1. 找到第 m-1 个节点
	for i := 1; i < m; i++ {
		pre = pre.Next
	}

	// 2. 反转 m 到 n 之间的节点，头插法
	start := pre.Next      // 第 m 个节点
	then := start.Next     // 第 m+1 个节点

	for i := 0; i < n-m; i++ {
		start.Next = then.Next
		then.Next = pre.Next
		pre.Next = then
		then = start.Next
	}

	return dummy.Next
}


*/
