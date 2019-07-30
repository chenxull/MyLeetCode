package problem231

/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two
 */
func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}

	is := n & (n - 1)
	return is == 0
}


