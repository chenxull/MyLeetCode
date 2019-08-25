package problem53

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */
/*

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
*/
//  寻找最大和的连续子数组
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 动态规划，要求最大值，使用 dp[i]保存当前位置和的最大值
	//  初始阶段 dp 中的所有值为 0，对于位置 i，关键看 dp[i-1]的值是正还是负
	//  dp[i] 表示到i位置，目前连续子数组的和最大是多少,所以在遍历的过程中，使用 Max 变量保存出现过的最大值
	dp := make([]int, len(nums))
	Max := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		// 每个位置，都要计算一下当前的最大和是多少
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
