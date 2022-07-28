package leetcode

// 判断链表是否有环
func detectCycle(head *ListNode) bool {
	return true
}

// 快慢指针
func detectCycleV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	for slow, fast := head, head; fast != nil; {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			node := head
			for node != slow {
				node = node.Next
				slow = slow.Next
			}
			return node
		}
	}
	return nil
}

// 暴力解答
// 用 hash 表存链表的内存地址值
func detectCycleV1(head *ListNode) *ListNode {
	m := map[*ListNode]byte{}
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = byte(1)
		head = head.Next
	}
	return nil
}
