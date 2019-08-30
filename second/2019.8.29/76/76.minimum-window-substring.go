package problem76

/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */
func minWindow(s string, t string) string {
	need := [256]int{}
	have := [256]int{}

	// 建立 hash 表
	for i := range t {
		need[t[i]]++
	}

	min := len(s) + 1
	size, totle := len(s), len(t)
	res := ""
	for i, j, count := 0, 0, 0; j < size; j++ {
		if have[s[j]] < need[s[j]] {
			// s[i] 位置处，有 need 中需要的元素
			count++
		}
		have[s[j]]++
		for i <= j && have[s[i]] > need[s[i]] {
			have[s[i]]--
			i++
		}

		width := j - i + 1
		if count == totle && min > width {
			min = width
			res = s[i : j+1]
		}
	}
	return res
}
