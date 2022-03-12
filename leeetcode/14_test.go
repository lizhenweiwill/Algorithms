package leetcode

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	var source []string
	source = []string{"flower", "flow", "flight"}
	t.Log(longestCommonPrefix(source))
	t.Log(longestCommonPrefix(source) == "fl")
	source = []string{"dog", "racecar", "car"}
	t.Log(longestCommonPrefix(source))
	t.Log(longestCommonPrefix(source) == "")
	source = []string{"a"}
	t.Log(longestCommonPrefix(source))
	t.Log(longestCommonPrefix(source) == "a")
	source = []string{"flower", "flower", "flower", "flower"}
	t.Log(longestCommonPrefix(source))
	t.Log(longestCommonPrefix(source) == "flower")
}
