package problem521

/*
 * @lc app=leetcode id=521 lang=golang
 *
 * [521] Longest Uncommon Subsequence I
 */
func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	return max(len(a), len(b))
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
