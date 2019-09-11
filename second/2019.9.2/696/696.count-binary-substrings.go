package problem696

/*
 * @lc app=leetcode id=696 lang=golang
 *
 * [696] Count Binary Substrings
 */
//  因为 01 是交替出现的，只要统计每种类型出现的次数即可
func countBinarySubstrings(s string) int {
	Rune := []rune(s)
	groups := make([]int, len(s))

	groups[0] = 1
	index := 0
	for i := 1; i < len(s); i++ {
		if Rune[i-1] != Rune[i] {
			index++
			groups[index] = 1
		} else {
			groups[index]++
		}
	}

	res := 0
	for i := 1; i < len(s); i++ {
		res += min(groups[i], groups[i-1])
	}
	return res

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
