package leetcode

// k 个一组反转节点
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 特殊情况返回
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	// 新建 hair 节点并将地址记录到 pre 中
	hair := &ListNode{Next: head}
	pre := hair
	// head 依次指向每组的头节点
	for head != nil {
		// tail 依次指向每组的尾节点
		tail := pre
		// tail 向后移动 k 步
		for i := 0; i < k; i++ {
			tail = tail.Next
			// 剩余节点数不足 k 个
			if tail == nil {
				// 返回 head 节点
				return hair.Next
			}
		}
		// 用 next 记住 tail 的下一个节点
		next := tail.Next
		// 将 head 和 tail 范围内的链表进行反转
		head, tail = reverse(head, tail)
		// pre 的 Next 指向 新 head
		pre.Next = head
		// tail 的 Next 指向 原 next
		tail.Next = next
		// pre 调整为下组头节点的前一个节点（ 即上组的 tail ）
		pre = tail
		// head 调整为新组的头节点
		head = tail.Next
	}
	// head 为 null 说明剩下的节点数
	return hair.Next
}

func reverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next
	node := head
	for pre != tail {
		next := node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return tail, head
}
