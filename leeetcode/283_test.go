package leetcode

import "testing"

func TestMoveZeroesV1(t *testing.T) {
	num := []int{0, 1, 0, 3, 12}
	t.Log(num)
	moveZeroesV1(num)
	t.Log(num)
}
