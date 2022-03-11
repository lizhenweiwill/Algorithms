package leetcode

import "testing"

func TestEvalRPN(t *testing.T) {
	var tokens []string
	tokens = []string{"2", "1", "+", "3", "*"}
	t.Log(evalRPN(tokens) == 9)
	tokens = []string{"4", "13", "5", "/", "+"}
	t.Log(evalRPN(tokens) == 6)
	tokens = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	t.Log(evalRPN(tokens) == 22)
}
