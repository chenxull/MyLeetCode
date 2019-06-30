package problem387

/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 *
 * https://leetcode.com/problems/first-unique-character-in-a-string/description/
 *
 * algorithms
 * Easy (48.98%)
 * Total Accepted:    251K
 * Total Submissions: 505.6K
 * Testcase Example:  '"leetcode"'
 *
 *
 * Given a string, find the first non-repeating character in it and return it's
 * index. If it doesn't exist, return -1.
 *
 * Examples:
 *
 * s = "leetcode"
 * return 0.
 *
 * s = "loveleetcode",
 * return 2.
 *
 *
 *
 *
 * Note: You may assume the string contain only lowercase letters.
 *
 */
func firstUniqChar(s string) int {
	ans := make([]int, 26)
	for _, v := range s {
		ans[v-'a']++
	}
	for i, v := range s {
		if ans[v-'a'] == 1 {
			return i
		}
	}
	return -1
}
