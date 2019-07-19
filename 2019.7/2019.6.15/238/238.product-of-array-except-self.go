package problem238

/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */
func productExceptSelf(nums []int) []int {

	length := len(nums)
	ans := make([]int, len(nums))

	left := 1
	for i := 0; i < length; i++ {
		ans[i] = left
		left *= nums[i]
	}

	right := 1
	for i := length - 1; i >= 0; i-- {
		ans[i] *= right
		right *= nums[i]
	}
	return ans
}
