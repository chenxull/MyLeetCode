package problem162

/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 */
func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1

	for low < high {
		mid := (low + high) / 2
		if nums[mid] < nums[mid+1] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
