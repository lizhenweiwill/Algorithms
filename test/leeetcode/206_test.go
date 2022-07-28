package leetcode

import "testing"

func TestReverseList(t *testing.T) {
	n5 := ListNode{5, nil}
	n4 := ListNode{4, &n5}
	n3 := ListNode{3, &n4}
	n2 := ListNode{2, &n3}
	n1 := ListNode{1, &n2}
	printNodeLise(&n1)
	node := reverseList(&n1)
	printNodeLise(node)
}
