package problem125

import "strings"

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1

	for i < j {
		for i < j && !ischar(s[i]) {
			i++
		}
		for i < j && !ischar(s[j]) {
			j--
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func ischar(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}
