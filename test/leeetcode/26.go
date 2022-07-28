package leetcode

// num 是 升序排列 的数组
// 删除重复元素，使得每个元素只出现一次，并保证相对顺序一致
func removeDuplicates(nums []int) int {
	return removeDuplicatesV2(nums)
}

// 双指针
// 非常类似于快慢指针
// 因为数组有序，所以和前一个数比较就知道是否为新的较大值
func removeDuplicatesV2(nums []int) int {
	n := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			nums[n] = nums[i]
			n++
		}
	}
	return n
}

// 快慢指针
// 快指针每次走一步
// 慢指针满足条件后走一步
func removeDuplicatesV1(nums []int) int {
	var slow = 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] > nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
