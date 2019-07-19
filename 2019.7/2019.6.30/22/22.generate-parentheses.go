package problem22

/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */
func generateParenthesis(n int) []string {

	res := make([]string, 0, n*n)
	bytes := make([]byte, 2*n)
	dfs(n, n, 0, bytes, &res)
	return res
}

func dfs(left, right, index int, bytes []byte, res *[]string) {
	// 返回条件
	if left == 0 && right == 0 {
		*res = append(*res, string(bytes))
		return
	}

	//	 当 left >0 时就不停的加入 stack 中
	if left > 0 {
		bytes[index] = '('
		dfs(left-1, right, index+1, bytes, res)
	}

	//	 当 right > 0 时，不停的将 ） 加入 stack 中
	if right > 0 && left < right {
		bytes[index] = ')'
		dfs(left, right-1, index+1, bytes, res)
	}
}
