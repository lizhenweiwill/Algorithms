package msb

// BubbleSort 冒泡排序
// 数组为 nil 无需排序
// 数组长度 <2 天然有序
// 在 0 ~ ( len - 1 ) - 0 之间
//     arr[i - 1] 和 arr[i] 比较，谁的值大谁往后，直到最大值到最后
// 在 0 ~ ( len - 1 ) - 1 之间
//     arr[i - 1] 和 arr[i] 比较，谁的值大谁往后，直到最大值到最后
// 在 0 ~ ( len - 1 ) - 2 之间
//     arr[i - 1] 和 arr[i] 比较，谁的值大谁往后，直到最大值到最后
func BubbleSort(arr []int) {
	// 边界检测
	if arr == nil {
		return
	}
	// 天然有序
	aLen := len(arr)
	if aLen < 2 {
		return
	}
	// 冒泡排序
	// 外层循环控制数组的结束位置
	for end := aLen - 1; end >= 0; end-- {
		// 内层循环控制相邻的两个元素的比较
		// 使用 arr[i-1] 和 arr[i] 的方式保证不越界
		for current := 1; current <= end; current++ {
			// 谁大谁往后挪
			if arr[current-1] > arr[current] {
				arr[current-1], arr[current] = arr[current], arr[current-1]
			}
		}
	}
}
