package problem77

/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */
func removeDuplicates(nums []int) int {
	size := len(nums)
	i := 0
	for j := 0; j < size; j++ {
		if i < 2 || nums[j] != nums[i-2] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
