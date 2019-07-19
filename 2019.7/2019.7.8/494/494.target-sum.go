package problem494

/*
 * @lc app=leetcode id=494 lang=golang
 *
 * [494] Target Sum
 */
//   难
func findTargetSumWays(nums []int, S int) int {
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum < S {
		return 0
	}

	//  必须为偶数，公式推导
	if (sum+S)%2 == 1 {
		return 0
	}

	return calSumWays(nums, (sum+S)/2)
}

//  将问题转化为求 nums 中有多少个子集和为(target +sum )/2 个数
func calSumWays(nums []int, target int) int {
	//  dp[i] 表示nums 中和为 i 的子集个数
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[target]
}
