package leetcode

import "strconv"

// 逆波兰表达式求值
func evalRPN(tokens []string) int {
	// 模拟数组栈
	var stack []int
	// 简单工厂
	m := map[string]func(a int, b int) int{
		"+": func(a int, b int) int { return a + b },
		"-": func(a int, b int) int { return a - b },
		"*": func(a int, b int) int { return a * b },
		"/": func(a int, b int) int { return a / b },
	}
	// 遍历表达式
	for _, v := range tokens {
		if f, ok := m[v]; ok {
			// 操作符
			n := len(stack)
			stack[n-2] = f(stack[n-2], stack[n-1])
			stack = stack[:n-1]
		} else {
			// 数字
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		}
	}
	return stack[0]
}
