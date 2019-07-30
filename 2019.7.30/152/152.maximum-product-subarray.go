package problem152

/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

//  不会写
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	r := nums[0]
	imin := r
	imax := r
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			temp := imin
			imin = imax
			imax = temp
		}

		imax = max(nums[i], imax*nums[i])
		imin = min(nums[i], imin*nums[i])

		r = max(r, imax)
	}
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
