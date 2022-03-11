package leetcode

// 判断链表是否有环
func hasCycle(head *ListNode) bool {
	return true
}

// 快慢指针
func hasCycleV2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	for slow, fast := head, head.Next; fast != slow; {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// 暴力解答
// 用 hash 表存链表的内存地址值
func hasCycleV1(head *ListNode) bool {
	m := map[*ListNode]byte{}
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = byte(1)
		head = head.Next
	}
	return false
}
