package problem342

/*
 * @lc app=leetcode id=342 lang=golang
 *
 * [342] Power of Four
 */
func isPowerOfFour(num int) bool {
	if num > 1 {
		for num%4 == 0 {
			num /= 4
		}
	}
	return num == 1
}
