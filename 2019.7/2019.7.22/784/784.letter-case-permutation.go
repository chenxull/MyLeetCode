package problem784

/*
 * @lc app=leetcode id=784 lang=golang
 *
 * [784] Letter Case Permutation
 */
func letterCasePermutation(S string) []string {
	if len(S) == 0 {
		return []string{}
	}

	res := []string{}
	backtrack(S, 0, "", &res)
	return res
}

//  直接在 S 上进行操作
func backtrack(S string, index int, cur string, res *[]string) {
	if index == len(S) {
		*res = append(*res, cur)
		return
	}

	for i := 0; i < len(S); i++ {
		if '0' <= S[i] && S[i] <= '9' {

		}

	}
}
