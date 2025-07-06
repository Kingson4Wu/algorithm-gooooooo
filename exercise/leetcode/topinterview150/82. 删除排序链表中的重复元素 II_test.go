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
/**
自己做的，用一个完全新的指针遍历（链表而不是数组，构建子链不增加空间复杂度。勇敢地构造子链即可）
每个结点前后对比， 符合条件就往新链表追加即可，记得结尾置空，不然旧结点仍存在
*/
func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	var newHead, pre, newHeadNext *ListNode
	// 每个结点前后对比
	for head != nil {
		if !same(pre, head) {
			if newHead == nil {
				newHead = head
				newHeadNext = head
			} else {
				newHeadNext.Next = head
				newHeadNext = newHeadNext.Next
			}
		}
		pre = head
		head = head.Next
	}
	//记得结尾置空
	if newHeadNext != nil {
		newHeadNext.Next = nil
	}

	return newHead
}

func same(pre, head *ListNode) bool {
	if pre != nil && head.Val == pre.Val {
		return true
	}
	if head.Next != nil && head.Val == head.Next.Val {
		return true
	}
	return false
}

/**
解答错误
96 / 166 个通过的测试用例

官方题解
输入
head =
[1,2,2]

添加到测试用例
输出
[1,2,2]
预期结果
[1]


执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.54
MB
击败
97.73%

*/
