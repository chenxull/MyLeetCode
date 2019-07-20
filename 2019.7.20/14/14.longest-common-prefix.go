package problem14

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	short := shorest(strs)

	// 以最小的字符串作为基准
	for i, v := range short {
		// 对所有的字符串进行对比
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != byte(v) {
				return short[:i]
			}
		}
	}
	return short
}

//  找到最小的
func shorest(strs []string) string {
	res := strs[0]

	for _, s := range strs {
		if len(res) > len(s) {
			res = s
		}
	}
	return res
}
