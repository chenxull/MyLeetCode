package problem153

/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */

//
func findMin(nums []int) int {
	var pivot int
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			pivot = i
		}
	}

	swap(0, pivot-1, nums)
	swap(pivot, len(nums)-1, nums)
	swap(0, len(nums)-1, nums)
	return nums[0]
}

func swap(i, j int, nums []int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
