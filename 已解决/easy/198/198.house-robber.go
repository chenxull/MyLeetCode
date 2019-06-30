package problem198

/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

func rob(nums []int) int {
	ifRobberdPrevious := 0   // 抢劫了前一个
	ifDidnotRobPrevious := 0 // 没有抢劫前一个

	// 遍历整个数组，维护二个变量：1)如果抢劫当前节点 2)不抢劫当前节点。
	for i := 0; i < len(nums); i++ {
		// 如果抢劫当前 cell，前一个 cell 不能抢劫，需要加上之前的结果
		currRober := ifDidnotRobPrevious + nums[i]

		//	 不抢劫当前 cell,值应该为抢劫或没有抢劫前一个 cell 中最大的值。
		currNotRober := max(ifRobberdPrevious, ifDidnotRobPrevious)

		// 更新值为下一次准备
		ifRobberdPrevious = currRober
		ifDidnotRobPrevious = currNotRober
	}
	return max(ifDidnotRobPrevious, ifRobberdPrevious)
	//if len(nums) == 0 {
	//	return 0
	//}
	//dp := make([]int, len(nums)+2)
	//
	//dp[0] = 0
	//dp[1] = 0
	//for i := 2; i < len(nums)+2; i++ {
	//	dp[i] = max(nums[i-2]+dp[i-2], dp[i-1])
	//}
	//return dp[len(nums)+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
