package msb

// InsertSort 插入排序
// 数组为 nil 无需排序
// 数组长度 <2 天然有序
// 类似于打牌时的排序方法
// 在 0 ~ 1 的范围上，0 ~ 0 的范围有序
//     新元素 arr[1] 和 pre ( arr[0] ) 比较，谁小谁向前移动
// 在 0 ~ 2 的范围上，0 ~ 1 的范围有序
//	   新元素 arr[2] 和 pre ( arr[0]，arr[1] ) 比较，谁小谁向前移动
// 在 0 ~ 3 的范围上，0 ~ 2 的范围有序
//     新元素 arr[3] 和 pre ( arr[0]，arr[1]，arr[2] ) 比较，谁小谁向前移动
func InsertSort(arr []int) {
	// 边界检测
	if arr == nil {
		return
	}
	// 天然有序
	aLen := len(arr)
	if aLen < 2 {
		return
	}
	// 插入排序
	// 外层循环控制 end 的位置
	// 最新加入的数肯定在 end 位置
	for current := 0; current < aLen; current++ {
		// 内层循环控制 pre 的位置
		for pre := current - 1; pre >= 0 && arr[pre] > arr[pre+1]; pre-- {
			// arr[pre] 和 arr[pre+1] 谁小谁向前移动
			arr[pre], arr[pre+1] = arr[pre+1], arr[pre]
		}
	}
}
