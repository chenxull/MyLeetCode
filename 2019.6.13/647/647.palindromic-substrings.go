package problem647

/*
 * @lc app=leetcode id=647 lang=golang
 *
 * [647] Palindromic Substrings
 */
func countSubstrings(s string) int {

	cnt := 0
	for i := 0; i < len(s); i++ {
		cnt += extendPalindrome(s, i, i)
		cnt += extendPalindrome(s, i, i+1)
	}
	return cnt

}

func extendPalindrome(s string, left, right int) int {
	res := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		res++
		left--
		right++
	}
	return res
}
