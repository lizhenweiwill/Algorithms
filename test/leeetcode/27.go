package leetcode

// 移除元素
// 给定数组 num 和元素 val，返回删除 val 后数组的长度
func removeElement(nums []int, val int) int {
	return removeElementV2(nums, val)
}

func removeElementV2(nums []int, val int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] == val {
			nums[l] = nums[r]
			r--
		} else {
			l++
		}
	}
	return l
}

func removeElementV1(nums []int, val int) int {
	n := 0
	for _, v := range nums {
		if v != val {
			nums[n] = v
			n++
		}
	}
	return n
}
