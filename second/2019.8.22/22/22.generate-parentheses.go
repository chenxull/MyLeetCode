package main

import "fmt"

/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 [
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/
//  使用回溯法
func generateParenthesis(n int) []string {
	res := []string{}
	cur := ""
	backtrack(cur, 0, 0, n, &res)
	return res
}

func backtrack(cur string, left, right, n int, res *[]string) {
	if len(cur) == 2*n {
		*res = append(*res, cur)
		return
	}

	if left < n {
		backtrack(cur+"(", left+1, right, n, res)
	}
	if right < left {
		backtrack(cur+")", left, right+1, n, res)
	}
}

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))
}
