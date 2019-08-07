package problem5

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

//  abba 00  01 12 03  -1 4
var lo, maxLen int

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
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
		j++
		i--
	}

	// 因为 i 和 j 都多进行了一次计算，所以在合理比较大小的时候，需要减一
	// 在求最小值坐标的时候，加 1 即可 
	if maxLen < j-i-1 {
		maxLen = j - i - 1
		lo = i + 1
	}

}
