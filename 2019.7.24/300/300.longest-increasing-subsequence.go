package problem300

/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */
//  最长上升子序列问题是在一个无序的给定序列中找到一个尽可能长的由低到高排列的子序列，
// 这种子序列不一定是连续的或者唯一的。
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	seq := make([]int, len(nums))

	seq[0] = 1
	m := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				seq[i] = max(seq[i], seq[j])
			}
		}
		seq[i]++
		m = max(m, seq[i])
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
