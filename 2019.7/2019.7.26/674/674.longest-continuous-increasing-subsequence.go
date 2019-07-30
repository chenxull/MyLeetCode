package problem674

/*
 * @lc app=leetcode id=674 lang=golang
 *
 * [674] Longest Continuous Increasing Subsequence
 */
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := 1
	i, j := 0, 1
	for j < len(nums) {
		// 寻找连续子串
		for j < len(nums) && nums[j-1] < nums[j] {
			j++
		}

		if res < j-i {
			res = j - i
		}
		i = j
		j++
	}
	return res
}
