package problem7

import "math"

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */
func reverse(x int) int {
	sign := 1

	if x < 0 {
		sign = -1
		x = -1 * x
	}

	res := 0
	for x > 0 {
		temp := x % 10
		res = 10*res + temp
		x = x / 10
	}

	res = sign * res

	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}

	return res
}
