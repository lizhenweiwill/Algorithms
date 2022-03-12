package leetcode

import "testing"

func TestRemoveElementV1(t *testing.T) {
	source := []int{0, 1, 2, 2, 3, 0, 4, 2}
	v1 := removeElementV1(source, 2)
	t.Log(v1)
}

func TestRemoveElementV2(t *testing.T) {
	source := []int{0, 1, 2, 2, 3, 0, 4, 2}
	v1 := removeElementV2(source, 2)
	t.Log(v1)
}
