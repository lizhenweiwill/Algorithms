package leetcode

// 实现 strStr()
// 相当于 Java 中的 indexOf()
func strStr(haystack string, needle string) int {
	return indexOfV1(haystack, needle)
}

func indexOfV1(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	m := len(needle)
	for i := 0; i <= len(haystack)-m; i++ {
		s := haystack[i : i+m]
		if s == needle {
			return i
		}
	}
	return -1
}

// KMS 算法
func indexOfV2(haystack string, needle string) int {
	return 0
}

// Sunday 算法
func indexOfV3(haystack string, needle string) int {
	return 0
}
