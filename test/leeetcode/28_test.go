package leetcode

import "testing"

func TestStrStr(t *testing.T) {
	t.Log(strStr("hello", "el") == 1)
	t.Log(strStr("hello", "ll") == 2)
	t.Log(strStr("aaaaa", "bba") == -1)
	t.Log(strStr("", "") == 0)
	t.Log(strStr("a", "a") == 0)
	t.Log(strStr("abc", "c") == 2)
}
