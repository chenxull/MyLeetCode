package problem198

/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums)+1)
	//  注意：这里 dp[0] 设置为 0 就是表示不偷第一家的
	// dp[1] 为偷第一家的
	dp[0] = 0
	dp[1] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i+1] = max(dp[i], dp[i-1]+nums[i])
	}
	return dp[len(nums)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
