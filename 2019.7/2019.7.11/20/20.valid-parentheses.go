package problem20

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */
//
func isValid(s string) bool {
	size := len(s)

	stack := make([]byte, size)
	top := 0
	// 这里利用到了 ascii 码的性质，直接将对应的括号的 ascii 码值存储在 stack 中
	// 比较的时候直接比较即可，无需做进一步判断
	for i := 0; i < size; i++ {
		c := s[i]
		switch c {
		case '(':
			stack[top] = c + 1 // c+1 = )
			top++
		case '{', '[':
			stack[top] = c + 2 //c+2 = } ]
			top++
		case ')', ']', '}':
			if top > 0 && stack[top-1] == c {
				top--
			} else {
				return false
			}

		}
	}
	return top == 0
}
