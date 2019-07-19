package problem172

/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */
func trailingZeroes(n int) int {
	res := 0
	for n >= 5 {
		n /= 5
		res += n
	}
	return res
}
