package hot100

/**
时间
116 ms
击败
16.55%
内存
5.1 MB
击败
60.64%

直接遍历调合并两个的方法即可
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	root := lists[0]
	for i := 1; i < len(lists); i++ {
		root = mergeTwoLists(root, lists[i])
	}
	return root
}
