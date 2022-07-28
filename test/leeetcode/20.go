package leetcode

// 有效的括号
// 给定只包括 () {} [] 的字符串，判断字符串是否有效
// 左括号必须按照顺序和相同类型的右括号匹配才能闭合
func isValid(s string) bool {
	//return isValid(s)
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	bh := ByteHeap{}
	for _, v := range s {
		b := byte(v)
		if r, ok := pairs[b]; ok {
			// 右括号
			if bh.Len() == 0 {
				return false
			}
			pop := bh.Pop()
			if r != pop {
				return false
			}
		} else {
			// 左括号
			bh.Push(b)
		}
	}
	return bh.Len() == 0
}

type ByteHeap []byte

func (b ByteHeap) Len() int {
	return len(b)
}

func (b ByteHeap) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b ByteHeap) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b *ByteHeap) Push(x interface{}) {
	*b = append(*b, x.(byte))
}

func (b *ByteHeap) Pop() interface{} {
	old := *b
	n := len(old)
	x := old[n-1]
	*b = old[0 : n-1]
	return x
}

func isValidV1(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for _, v := range s {
		if b, ok := pairs[byte(v)]; ok {
			// 右括号
			n := len(stack)
			if n == 0 {
				return false
			}
			i := n - 1
			if stack[i] != b {
				return false
			}
			stack = stack[:i]
		} else {
			// 左括号
			stack = append(stack, byte(v))
		}
	}
	return len(stack) == 0
}
