package problem290

import "strings"

/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 */
func wordPattern(pattern string, str string) bool {
	pat := strings.Split(pattern, "")
	sr := strings.Split(str, " ")

	if len(pat) != len(sr) {
		return false
	}

	return isPattern(pat, sr) && isPattern(sr, pat)
}

func isPattern(n, m []string) bool {
	size := len(n)
	pattern := make(map[string]string, size)

	for i := 0; i < size; i++ {

		// 只有当 pattern 中存在此元素时，ok 才为 true
		if w, ok := pattern[n[i]]; ok {
			if w != m[i] {
				return false
			}
		} else {
			// 构建 map
			pattern[n[i]] = m[i]
		}
	}
	return true
}
