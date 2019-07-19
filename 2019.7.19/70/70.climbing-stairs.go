package problem70

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */
func climbStairs(n int) int {
	if n < 2 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
