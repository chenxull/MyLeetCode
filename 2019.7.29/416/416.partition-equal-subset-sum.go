package problem416

/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 */
//  dp[i]保存的是所有可能出现的和， dp[i] = ture 表示 和为 i 存在
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	dp := make([]bool, sum+1)
	dp[0] = true
	for _, num := range nums {
		for i := sum; i >= 0; i-- {
			if dp[i] {
				dp[i+num] = true
			}
		}
		if dp[sum/2] {
			return true
		}
	}
	return false
}
