package week01

import (
	"reflect"
	"testing"
)

func TestMergeTwoListsV1(t *testing.T) {

}

func TestMergeTwoLists(t *testing.T) {
	// 创建六个节点
	n1 := ListNode{Val: 1, Next: nil}
	n2 := ListNode{Val: 2, Next: nil}
	n3 := ListNode{Val: 4, Next: nil}
	n4 := ListNode{Val: 1, Next: nil}
	n5 := ListNode{Val: 3, Next: nil}
	n6 := ListNode{Val: 4, Next: nil}
	// 创建链表 1
	n1.Next = &n2
	n2.Next = &n3
	list1 := &n1
	// 创建链表 2
	n4.Next = &n5
	n5.Next = &n6
	list2 := &n4
	// 返回合并后的链表
	list := mergeTwoLists(list1, list2)
	var result []int
	for list != nil {
		result = append(result, list.Val)
		list = list.Next
	}
	target := []int{1, 1, 2, 3, 4, 4}
	// 判定是否相等
	if !reflect.DeepEqual(target, result) {
		t.Log("error")
	}
}
