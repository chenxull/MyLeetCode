package probelm673

/*
 * @lc app=leetcode id=673 lang=golang
 *
 * [673] Number of Longest Increasing Subsequence
 */

//   很复杂的一题
func findNumberOfLIS(nums []int) int {

	n := len(nums)
	res := 0
	maxLen := 0

	// length[i] == m 表示，以 a[i] 结尾的最长子序列的长度为 m
	length := make([]int, n)

	// count[i] == n 表示，以 a[i] 结尾的最长子序列的个数为 n
	count := make([]int, n)

	for j := 0; j < n; j++ {
		length[j] = 1
		count[j] = 1
		for i := 0; i < j; i++ {
			if nums[i] < nums[j] {
				// 找到相同长度 ,
				if length[j] == length[i]+1 {
					count[j] += count[i]
				} else if length[j] < length[i]+1 {
					length[j] = length[i] + 1
					count[j] = count[i]
				}
			}
		}

		// 对于每一次 j 的更新，都判断一次
		if maxLen == length[j] {
			res += count[j]
		} else if maxLen < length[j] {
			res = count[j]
			maxLen = length[j]
		}
	}
	return res

}
