package problem326

/*
 * @lc app=leetcode id=326 lang=golang
 *
 * [326] Power of Three
 */
func isPowerOfThree(n int) bool {
	if n > 1 {
		for n%3 == 0 {
			n = n / 3
		}
	}

	return n == 1
}
