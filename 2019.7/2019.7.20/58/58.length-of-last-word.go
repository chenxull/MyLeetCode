package problem58

/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */
//   有问题，不深究
func lengthOfLastWord(s string) int {
	// 是否要考虑
	if len(s) == 0 {
		return 0
	}
	size := len(s)
	res := 0

	for i := size - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if res != 0 {
				return res
			}
			continue
		}
		res++
	}

	return res
}
