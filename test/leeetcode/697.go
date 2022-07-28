package leetcode

type Info struct {
	start int
	end   int
	count int
}

// 第一步：先遍历数组，找到每个元素的起始位置和出现次数
// 第二步：找到最大的度（出现次数）对应的元素，可能有多个
// 第三步：num[end] - num[start] + 1 最小即为答案
func findShortestSubArray(nums []int) int {
	m := make(map[int]Info)
	for i, v := range nums {
		if info, ok := m[v]; ok {
			// map 存在该 key
			info.end = i
			info.count++
			m[v] = info
		} else {
			// map 不存在该 key
			m[v] = Info{i, i, 1}
		}
	}
	minLen := 0
	maxCount := 0
	for _, v := range m {
		if maxCount < v.count {
			maxCount = v.count
			minLen = v.end - v.start + 1
		} else if maxCount == v.count {
			n := v.end - v.start + 1
			if minLen > n {
				minLen = n
			}
		}
	}
	return minLen
}
