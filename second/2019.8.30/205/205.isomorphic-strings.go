package problem205

import "strings"

/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	bs := strings.Split(s, "")
	bt := strings.Split(t, "")

	if isSame(bs, bt) && isSame(bt, bs) {
		return true
	}
	return false
}

func isSame(bs, bt []string) bool {
	pattern := make(map[string]string, len(bs))
	for i := 0; i < len(bs); i++ {
		if value, ok := pattern[bs[i]]; ok {
			if value != bt[i] {
				return false
			}
		} else {
			pattern[bs[i]] = bt[i]

		}
	}
	return true
}
