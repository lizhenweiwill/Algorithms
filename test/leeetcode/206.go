package leetcode

// 已知单链表的头节点 head
// 反转链表
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode // nil
	node := head
	for node != nil {
		next := node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return pre
}
