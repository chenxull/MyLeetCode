package problem121

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minBuy := prices[0]
	maxprofit := 0
	for i := 0; i < len(prices); i++ {
		//  确保买入价格最低
		if prices[i] < minBuy {
			minBuy = prices[i]
		} else if prices[i]-minBuy > maxprofit {
			maxprofit = prices[i] - minBuy
		}
	}
	return maxprofit

}
