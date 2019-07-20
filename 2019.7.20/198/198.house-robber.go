package problem198

/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */
//   记忆化搜索算法
func rob(nums []int) int {
	//mem[i] 表示考虑去抢劫(i..n)所能获取的最大收益
	mem := make([]int, len(nums))
	return tryRob(nums, mem, 0)
}

// tryRob 抢劫[index...len(nums)]位置的房子
func tryRob(nums, mem []int, index int) int {
	if index > len(nums) {
		return 0
	}

	if mem[index] != 0 {
		return mem[index]
	}

	res := 0
	for i := index; i < len(nums); i++ {
		// 考虑强当前节点，以及后续的节点
		res = max(res, nums[i]+tryRob(nums, mem, i+2))
	}

	//  记录 index 位置，所能获得的最大收益
	mem[index] = res
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
