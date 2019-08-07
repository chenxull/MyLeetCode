package problem53

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */
//  使用动态规划，dp[i]表示 前 i 位中最大和
// 同时使用一个变量保存全局的最大值，最后返回这个最大值 
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 使用一个变量记录出现的最大值
	m := nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		m = max(m, dp[i])
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
