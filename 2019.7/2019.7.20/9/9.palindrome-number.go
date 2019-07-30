package problem9

import "strconv"

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)

	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
