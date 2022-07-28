package leetcode

import "testing"

func TestCalculate(t *testing.T) {
	var s string
	s = "1 + 1"
	t.Log(calculate(s) == 2)
	s = " 2-1 + 2"
	t.Log(calculate(s) == 3)
	s = "(1+(4+5+2)-3)+(6+8)"
	t.Log(calculate(s) == 23)
	s = "(3+(1+2)-3)+4"
	t.Log(calculate(s) == 7)
}
