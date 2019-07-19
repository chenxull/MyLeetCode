package problem279

import "math"

/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */

//  有待深入的理解
func numSquares(n int) int {
	perfects := []int{}
	for i := 1; i*i <= n; i++ {
		perfects = append(perfects, i*i)
	}

	// dp[i] 表示 the least number of perfect square numbers which sum to i
	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}

	for _, p := range perfects {
		for i := p; i < len(dp); i++ {
			if dp[i] > dp[i-p]+1 {
				// 因为 i = ( i - p ) + p，p 是 平方数
				// 所以 dp[i] = dp[i-p] + 1
				dp[i] = dp[i-p] + 1
			}
		}
	}

	return dp[n]
}
