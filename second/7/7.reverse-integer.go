package problem7

import "math"

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */
func reverse(x int) int {

	sig := 1
	if x < 0 {
		sig = -1
		x = -x
	}

	res := 0
	for x > 0 {
		temp := x % 10
		res = res*10 + temp
		x = x / 10
	}

	res = res * sig

	// 这道题唯一需要主要的就是 确保答案不要超出范围
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return res
}
