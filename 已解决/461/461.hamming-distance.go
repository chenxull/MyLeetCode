package problem461

/*
 * @lc app=leetcode id=461 lang=golang
 *
 * [461] Hamming Distance
 */
func hammingDistance(x int, y int) int {
	x ^= y
	res := 0
	if x > 0 {
		res += x & 1
		x = x >> 1
	}
	return res
}
