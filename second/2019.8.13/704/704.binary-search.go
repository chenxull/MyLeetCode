package problem704

/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// [l,r)
	l, r := 0, len(nums)

	for l < r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		}
		if nums[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	return -1
}
