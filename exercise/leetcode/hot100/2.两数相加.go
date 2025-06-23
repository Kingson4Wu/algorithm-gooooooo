package hot100

//leetcode/linked_list/2. 两数相加.go

// 自己新写的这么复杂，还不如以前的自己写的
// addTwoNumbers
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var l3 *ListNode
	var head *ListNode
	next := 0
	v1, v2 := 0, 0

	for l1 != nil || l2 != nil || next > 0 {

		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}

		sum := v1 + v2 + next
		remain := sum % 10
		next = sum / 10
		if remain > 0 || next > 0 {
			node := &ListNode{}
			node.Val = remain
			if head == nil {
				l3 = node
				head = node
			} else {
				head.Next = node
				head = node
			}
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		v1 = 0
		v2 = 0
	}
	return l3

}
