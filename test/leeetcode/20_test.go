package leetcode

import "testing"

func TestIsValid(t *testing.T) {
	var source string
	source = "()"
	t.Log(isValid(source) == true)
	source = "()[]{}"
	t.Log(isValid(source) == true)
	source = "(]"
	t.Log(isValid(source) == false)
	source = "([)]"
	t.Log(isValid(source) == false)
	source = "{[]}"
	t.Log(isValid(source) == true)
	source = "(("
	t.Log(isValid(source) == false)
	source = "){"
	t.Log(isValid(source) == false)
}
