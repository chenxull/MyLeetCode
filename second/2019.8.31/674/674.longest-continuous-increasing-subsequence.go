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
	max := 0
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			count++
		} else {
			max = ismax(max, count)
			count = 1

		}
	}
	max = ismax(count, max)
	return max
}

func ismax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
