package problem300

/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */
//  使用动态规划求解
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = 1
	m := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i]++
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
