package problem322

/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */
//   dp,从下到上。因为每个 amount 所需要的最少 coins 是确定的
func coinChange(coins []int, amount int) int {
	// dp[i] i 表示金额
	dp := make([]int, amount+1)
	dp[0] = 0
	// 计算出每个金额需要的最少硬币数量
	for i := 1; i < amount+1; i++ {
		//  每个位置都初始化为最大值
		dp[i] = amount + 1
		// 遍历每一金额的硬币
		for _, c := range coins {
			if c <= i && dp[i] > dp[i-c]+1 {
				dp[i] = dp[i-c] + 1
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
