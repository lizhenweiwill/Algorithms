package leetcode

import "testing"

func TestHasCycleV1(t *testing.T) {
	n7 := ListNode{7, nil}
	n6 := ListNode{6, &n7}
	n5 := ListNode{5, &n6}
	n4 := ListNode{4, &n5}
	n3 := ListNode{3, &n4}
	n2 := ListNode{2, &n3}
	n1 := ListNode{1, &n2}
	// 模拟环
	n7.Next = &n4
	//printNodeLise(&n1)
	t.Log(hasCycleV1(&n1))
}

func TestHasCycleV2(t *testing.T) {
	n7 := ListNode{7, nil}
	n6 := ListNode{6, &n7}
	n5 := ListNode{5, &n6}
	n4 := ListNode{4, &n5}
	n3 := ListNode{3, &n4}
	n2 := ListNode{2, &n3}
	n1 := ListNode{1, &n2}
	// 模拟环
	n7.Next = &n4
	//printNodeLise(&n1)
	t.Log(hasCycleV2(&n1))
}
