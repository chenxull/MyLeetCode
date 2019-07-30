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

	//  left,right 分别表示需要 (,)的数量
	left, right := 0, 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			right++ // 需要 ')'的数量+1
		} else if S[i] == ')' && right > 0 {
			right-- // 这个')'正好与之前的'(‘ 相匹配，需要的')'数量-1
		} else {
			left++
		}
	}
	return right + left

}
