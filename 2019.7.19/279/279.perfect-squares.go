package problem279

import "math"

/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */
func numSquares(n int) int {
	perfects := []int{}
	for i := 1; i*i <= n; i++ {
		perfects = append(perfects, i*i)
	}
	//  dp[i] 表示 i 可由多少了平方和的数字相加
	dp := make([]int, n+1)
	// 将其都置为最大值
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}

	// 动态规划的核心,遍历 perfect 中的所有元素，从低位开始判断
	for _, p := range perfects {
		// 获取 n 的长度,实际上 i 最大可以到达 n。这其实在计算
		// dp[n] = min(dp[n -i*i]+1)
		// 对于 n =13来说,dp[13] = min(dp[12]+1,dp[9]+1,dp[4]+1) 这种关系表达式求解下去
		for i := p; i < len(dp); i++ {
			// 只有当 更小的出现时，才赋值
			if dp[i] > dp[i-p]+1 {
				dp[i] = dp[i-p] + 1
			}
		}
	}
	return dp[n]

}
