package leetcode

import "testing"

func TestMergeV1(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	t.Log(nums1)
	t.Log(nums2)
	num := mergeV1(nums1, 3, nums2, 3)
	t.Log(num)
}

func TestMergeV2(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	t.Log(nums1)
	t.Log(nums2)
	num := mergeV2(nums1, 3, nums2, 3)
	t.Log(num)
}
