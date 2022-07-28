package msb

// SelectSort 选择排序
// 数组为 nil 无需排序
// 数组长度 <2 天然有序
// 从 0 位置开始遍历数组
//     从 1 开始找最小值的下标，和 0 位置交换，保证 0 位置最小
// 从 1 位置开始遍历数组
//     从 2 开始找最小值的下标，和 1 位置交换，保证 1 位置最小
// 从 2 位置开始遍历数组
//     从 3 开始找最小值的下标，和 2 位置交换，保证 2 位置最小
func SelectSort(arr []int) {
	// 边界条件
	if arr == nil {
		return
	}
	// 天然有序
	aLen := len(arr)
	if aLen < 2 {
		return
	}
	// 选择排序
	// 外层循环控制数组遍历
	for i := 0; i < aLen; i++ {
		// 内部变量记录最小值索引
		minValIndex := i
		// 内部循环寻找最小值
		for j := i + 1; j < aLen; j++ {
			// 如果当前下标的数小于最小值下标上的数
			// 更新最小值元素的下标为当前元素下标
			if arr[j] < arr[minValIndex] {
				minValIndex = j
			}
		}
		// 交换
		arr[i], arr[minValIndex] = arr[minValIndex], arr[i]
	}
}
