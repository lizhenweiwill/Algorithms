package leetcode

func printNodeLise(node *ListNode) {
	for node != nil {
		print(node.Val, "\t")
		node = node.Next
	}
	println()
}
