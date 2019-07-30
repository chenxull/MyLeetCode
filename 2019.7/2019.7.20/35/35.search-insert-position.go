package problem35

/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if target > nums[i] {
			continue
		}
		if target <= nums[i] {
			return i
		}
	}
	return len(nums)
}
