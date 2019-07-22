package problem22

/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := []string{}
	backtrack("", 0, 0, n, &res)
	return res
}

func backtrack(cur string, open, close, max int, res *[]string) {
	if len(cur) == 2*max {
		*res = append(*res, cur)
		return
	}

	if open < max {
		backtrack(cur+"(", open+1, close, max, res)
		// 每次从递归中出来，相当于 open-1 .同时，cur 是在递归的过程
		// 中加上"("的，退出时没有 (
	}
	if close < open {
		backtrack(cur+")", open, close+1, max, res)
	}

}
