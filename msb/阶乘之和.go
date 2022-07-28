package msb

// SumV0
// 暴力解答
func SumV0(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += factorialV0(i)
	}
	return sum
}

func factorialV0(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

// SumV1
// 递归
func SumV1(n int) int {
	sum := 0
	cur := 1
	for i := 1; i <= n; i++ {
		cur = cur * i
		sum += cur
	}
	return sum
}

func factorialV1(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return factorialV1(n-1) * n
}
