package problem28

/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */
//  使用滑动窗口的思想
func strStr(haystack string, needle string) int {

	hLen, nLen := len(haystack), len(needle)
	if hLen == 0 && nLen == 0 {
		return 0
	}

	if hLen == 0 {
		return -1
	}

	for i := 0; i <= hLen-nLen; i++ {
		if haystack[i:i+nLen] == needle {
			return i
		}
	}
	return -1
}
