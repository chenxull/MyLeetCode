package problem198

/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */
func rob(nums []int) int {
	var a, b int

	for i, v := range nums {
		if i%2 == 0 {
			a = max(a+v, b)
		} else {
			b = max(b+v, a)
		}
	}
	return max(a, b)
}

