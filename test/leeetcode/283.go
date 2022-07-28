package leetcode

// 将数组中的 0 全部移动到末尾
// 保持所有非 0 元素的相对位置
func moveZeroes(nums []int) {
	moveZeroesV1(nums)
}

// 快慢指针
func moveZeroesV1(nums []int) {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[n], nums[i] = nums[i], nums[n]
			n++
		}
	}
}
