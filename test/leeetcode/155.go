package leetcode

import "math"

// MinStack 最小栈
type MinStack struct {
	Item *Item
}

type Item struct {
	Val  int
	Min  int
	Next *Item
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.Item == nil {
		// 首次添加
		this.Item = &Item{
			Val:  val,
			Min:  val,
			Next: nil,
		}
	} else {
		// 非首次添加
		pre := this.Item
		item := Item{Val: val, Min: int(math.Min(float64(pre.Min), float64(val)))}
		item.Next = pre
		this.Item = &item
	}
}

func (this *MinStack) Pop() {
	next := this.Item.Next
	this.Item.Next = nil
	this.Item = next
}

func (this *MinStack) Top() int {
	return this.Item.Val
}

func (this *MinStack) GetMin() int {
	return this.Item.Min
}
