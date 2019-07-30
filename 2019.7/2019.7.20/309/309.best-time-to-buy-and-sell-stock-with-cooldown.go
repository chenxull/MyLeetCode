package problem309

import "math"

/*
 * @lc app=leetcode id=309 lang=golang
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 */
//   难
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	s0, s1, s2 := make([]int, len(prices)), make([]int, len(prices)), make([]int, len(prices))
	s0[0] = 0
	s1[0] = -prices[0]
	s2[0] = math.MinInt32

	for i := 1; i < len(prices); i++ {
		s0[i] = max(s0[i-1], s2[i-1])
		s1[i] = max(s1[i-1], s0[i-1]-prices[i])
		s2[i] = s1[i-1] + prices[i]
	}
	return max(s0[len(prices)-1], s2[len(prices)-1])
}

func max(a, b int) int {
	if a > b {
		return a

	}

	return b
}
