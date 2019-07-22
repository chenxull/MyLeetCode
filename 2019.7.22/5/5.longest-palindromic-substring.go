package problem5

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */
//  对于奇偶个数的字符串，分情况讨论
var lo, maxLen int

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	lo, maxLen = 0, 0

	for i := 0; i < len(s); i++ {
		palindrome(s, i, i)
		palindrome(s, i, i+1)

	}

	return s[lo : lo+maxLen]

}
func palindrome(s string, i, j int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}

	if maxLen < j-i-1 {
		lo = i + 1
		maxLen = j - i - 1
	}
}
