package problem209

/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */
//   滑动窗口 
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	// [i,j] 为滑动窗口
	minLen, i, j, sum := n+1, 0, 0, 0

	for j < n {
		sum += nums[j]
		j++
		for sum >= s {
			sum -= nums[i]
			i++
			if minLen > j-i+1 {
				minLen = j - i + 1
			}
		}
	}

	return minLen % (n + 1)
}
