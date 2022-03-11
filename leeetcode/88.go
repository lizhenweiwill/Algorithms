package leetcode

import "sort"

// 两个按 非递减顺序 排列的数组 num1 和 num2
// 合并 num2 到 num1 中，使合并后的数组同样为 非递减顺序
// num1 长度为 m+n，实际有 m 个元素
// num2 长度为 n，实际有 n 个元素
func merge(nums1 []int, m int, nums2 []int, n int) {
	var v int
	for p1, p2, i := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; i-- {
		if p1 < 0 {
			v = nums2[p2]
			p2--
		} else if p2 < 0 {
			v = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			v = nums1[p1]
			p1--
		} else {
			v = nums2[p2]
			p2--
		}
		nums1[i] = v
	}
}

// 暴力解
// 将 num2 拷贝到 num1 的末尾，然后调用系统排序方法
func mergeV1(nums1 []int, m int, nums2 []int, n int) []int {
	nums1 = append(nums1[0:m], nums2...)
	sort.Ints(nums1)
	return nums1
}

// 最优解
func mergeV2(nums1 []int, m int, nums2 []int, n int) []int {
	end1 := m - 1
	end2 := n - 1
	index := m + n - 1
	for end1 >= 0 || end2 >= 0 {
		if end1 < 0 {
			nums1[index] = nums2[end2]
			end2--
		} else if end2 < 0 {
			nums1[index] = nums1[end1]
			end1--
		} else if nums1[end1] > nums2[end2] {
			nums1[index] = nums1[end1]
			end1--
		} else {
			nums1[index] = nums2[end2]
			end2--
		}
		index--
	}
	return nums1
}
