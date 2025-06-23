package hot100

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	node := addTwoNumbers(l1, l2)

	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
