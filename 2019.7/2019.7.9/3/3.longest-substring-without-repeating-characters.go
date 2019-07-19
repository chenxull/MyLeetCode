package problem3

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
func lengthOfLongestSubstring(s string) int {
	i, j, n := 0, 0, len(s)
	res := 0
	hash := [256]int{}
	// 应为是求最大子字符串，所以可以直接以 j 为循环终止条件。
	for j < n {
		if hash[s[j]] == 0 {
			hash[s[j]]++
			j++
		} else {
			hash[s[i]]--
			i++
		}
		res = max(res, j-i)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
