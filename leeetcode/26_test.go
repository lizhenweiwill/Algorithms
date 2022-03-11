package leetcode

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	num := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	t.Log(num)
	n := removeDuplicates(num)
	t.Log(n)
	t.Log(num)
}

func TestRemoveDuplicatesV1(t *testing.T) {
	num := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	t.Log(num)
	n := removeDuplicatesV1(num)
	t.Log(n)
	t.Log(num)
}

func TestRemoveDuplicatesV2(t *testing.T) {
	num := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	t.Log(num)
	n := removeDuplicatesV2(num)
	t.Log(n)
	t.Log(num)
}
