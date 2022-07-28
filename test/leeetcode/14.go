package leetcode

func longestCommonPrefix(strings []string) string {
	if strings == nil || len(strings) == 0 {
		return ""
	}
	for i, v := range strings[0] {
		for _, str := range strings {
			if i == len(str) || str[i] != byte(v) {
				return strings[0][:i]
			}
		}
	}
	return strings[0]
}
