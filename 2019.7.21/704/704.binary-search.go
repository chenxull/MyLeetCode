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

	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else if nums[mid] > target {
			hi = mid
		} else {
			return mid
		}

	}
	return -1
}
