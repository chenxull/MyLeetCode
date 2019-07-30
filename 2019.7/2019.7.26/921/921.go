 package problem921

/*
 * @lc app=leetcode id=921 lang=golang
 *
 * [921] Minimum Add to Make Parentheses Valid
 */
func minAddToMakeValid1(S string) int {
	if len(S) == 0 {
		return 0
	}

	// ans 用来记录')'多的情况； balance 用来记录'('多的情况
	ans, balance := 0, 0

	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			balance++
		} else {
			balance--
		}

		// 突破平衡限制
		if balance == -1 {
			ans++
			balance++
		}
	}
	return ans + balance

}
