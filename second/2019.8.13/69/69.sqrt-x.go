package problem69

/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */
func mySqrt(x int) int {
	// [l,r) 处理区间
	l, r := 0, x+1
	for l < r {
		m := l + (r-l)/2
		if m*m > x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l - 1
}
