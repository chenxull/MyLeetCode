package problem921

/*
 * @lc app=leetcode id=921 lang=golang
 *
 * [921] Minimum Add to Make Parentheses Valid
 */
func minAddToMakeValid(S string) int {
	if len(S) == 0 {
		return 0
	}

	left, right := 0, 0

	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			left++
		} else {
			left--
		}
		//  ( 用完，有一个多的)括号
		if left == -1 {
			right++
			left++
		}
	}

	return left + right
}
