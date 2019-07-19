package problem371

/*
 * @lc app=leetcode id=371 lang=golang
 *
 * [371] Sum of Two Integers
 */
func getSum(a int, b int) int {
	for a != 0 {
		a, b = (a&b)<<1, a^b
	}
	return b
}
