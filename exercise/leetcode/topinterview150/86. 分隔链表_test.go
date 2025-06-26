package topinterview150

import "testing"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
自己写的，不过总感觉链表会写得很乱

因为是链表而不是数组，构建子链不增加空间复杂度。勇敢地构造子链即可，无需考虑节点交换。 ！！！！
*/
/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.06
MB
击败
63.43%
*/
/**
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

你应当 保留 两个分区中每个节点的初始相对位置。
*/
func partition(head *ListNode, x int) *ListNode {

	origin := head

	var pre *ListNode
	var small *ListNode

	for head != nil {

		if head.Val < x {
			if small != nil {
				if small == pre {
					small = head
				} else {
					temp := small.Next
					small.Next = head
					small = head
					pre.Next = head.Next
					head.Next = temp
					head = pre.Next
					continue
				}
			} else {
				if head == origin {
					small = head
				} else {
					pre.Next = head.Next
					head.Next = origin
					origin = head
					small = head
					head = pre.Next
					continue
				}
			}
		}

		pre = head
		head = head.Next
	}

	return origin
}

func TestPartition(t *testing.T) {
	partition(&ListNode{Val: 3, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}, 3)
}

/**
解答错误
138 / 168 个通过的测试用例

官方题解
输入
head =
[3,1,2]
x =
3

添加到测试用例
输出
[2,1,3]
预期结果
[1,2,3]
*/
