package problem53

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	Max := nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		Max = max(dp[i], Max)
	}
	return Max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
