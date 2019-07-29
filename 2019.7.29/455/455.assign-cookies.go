package problem455

import "sort"

/*
 * @lc app=leetcode id=455 lang=golang
 *
 * [455] Assign Cookies
 */
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	res := 0
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			res++
			i++
		}
		j++
	}
	return res
}
