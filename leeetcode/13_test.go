package leetcode

import "testing"

func TestRomanToInt(t *testing.T) {
	var source string
	source = "II"
	t.Log(romanToInt(source) == 2)
	source = "IV"
	t.Log(romanToInt(source) == 4)
	source = "IX"
	t.Log(romanToInt(source) == 9)
	source = "LVIII"
	t.Log(romanToInt(source) == 58)
	source = "MCMXCIV"
	t.Log(romanToInt(source) == 1994)
}
