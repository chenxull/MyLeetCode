package problem409

/*
 * @lc app=leetcode id=409 lang=golang
 *
 * [409] Longest Palindrome
 */

func longestPalindrome(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	odd := 0
	hashmap := make(map[byte]int, len(s))
	for i := range s {
		hashmap[s[i]]++
	}

	for _, v := range hashmap {
		if v%2 == 0 {
			res += v
		} else {
			//  对于多个奇数，只需要减去一个即可
			res += v - 1
			odd = 1
		}
	}
	return res + odd
}
