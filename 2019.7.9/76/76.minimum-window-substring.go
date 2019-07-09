package problem76

/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */
func minWindow(s string, t string) string {
	need := [256]int{}
	have := [256]int{}
	//  建立目标元素的 hash 表
	for i := range t {
		need[t[i]]++
	}
	min := len(s) + 1
	size, total := len(s), len(t)
	res := ""
	for i, j, count := 0, 0, 0; j < size; j++ {
		if have[s[j]] < need[s[j]] {
			//  这个字符是需要的，且不存在于滑动窗口中
			count++
		}
		// 将窗口最右端的未添加的元素，添加进窗口
		have[s[j]]++

		// 在确保窗口不丢失所需的字母情况下，i 要尽可能的大
		for i <= j && have[s[i]] > need[s[i]] {
			have[s[i]]--
			i++
		}

		width := j - i + 1
		if count == total && min > width {
			min = width
			res = s[i : j+1]
		}
	}
	return res
}
