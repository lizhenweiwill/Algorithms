package leetcode

import "testing"

func TestMinStack(t *testing.T) {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	t.Log(obj.GetMin() == -3)
	obj.Pop()
	t.Log(obj.Top() == 0)
	t.Log(obj.GetMin() == -2)
}

func TestMinStack2(t *testing.T) {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-1)
	t.Log(obj.GetMin() == -2)
	t.Log(obj.Top() == -1)
	obj.Pop()
	t.Log(obj.GetMin() == -2)
}
