package probelm279

import "math"

/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */
func numSquares(n int) int {
	perfect := []int{}

	for i := 1; i*i < n; i++ {
		perfect = append(perfect, i*i)
	}

	// dp[n] 表示 n 位置所需要的 perfect num 个数
	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt64
	}

	for _, p := range perfect {
		for i := p; i < len(dp); i++ {
			if dp[i] > dp[i-p]+1 {
				dp[i] = dp[i-p] + 1
			}
		}
	}

	return dp[n]
}
