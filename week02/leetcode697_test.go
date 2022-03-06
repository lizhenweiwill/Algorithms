package week02

import "testing"

func TestFindShortestSubArray(t *testing.T) {
	source := []int{1}
	n := findShortestSubArray(source)
	println(n)
}
