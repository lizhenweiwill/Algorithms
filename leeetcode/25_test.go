package leetcode

import "testing"

func Test(t *testing.T) {
	n7 := ListNode{7, nil}
	n6 := ListNode{6, &n7}
	n5 := ListNode{5, &n6}
	n4 := ListNode{4, &n5}
	n3 := ListNode{3, &n4}
	n2 := ListNode{2, &n3}
	n1 := ListNode{1, &n2}
	printNodeLise(&n1)
	node := reverseKGroup(&n1, 2)
	printNodeLise(node)
}

func TestReverse(t *testing.T) {
	n7 := ListNode{7, nil}
	n6 := ListNode{6, &n7}
	n5 := ListNode{5, &n6}
	n4 := ListNode{4, &n5}
	n3 := ListNode{3, &n4}
	n2 := ListNode{2, &n3}
	n1 := ListNode{1, &n2}
	printNodeLise(&n1)
	head, _ := reverse(&n1, &n7)
	printNodeLise(head)
}
