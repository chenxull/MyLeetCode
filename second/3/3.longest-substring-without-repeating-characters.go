package problem3

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
func lengthOfLongestSubstring(s string) int {
	// 使用滑动窗口的思想来解决，最长子连续不重复数组问题

	// i用来保存窗口左边的数据，j 用来保存窗口右边的元素
	i, j, size := 0, 0, len(s)
	res := 0
	hash := [256]int{}
	for j < size {
		if hash[s[j]] == 0 {
			hash[s[j]]++
			j++
		} else { // 出现重复元素，窗口最左边的元素 数量减一并且向右移动一位
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
