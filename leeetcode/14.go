package leetcode

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	for i, v := range strs[0] {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != byte(v) {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
