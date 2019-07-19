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

	//  因为是求最小的长度，所以左边的索引需要不停的计算
	for i < n {

		if j < n && sum < s {
			sum += nums[j]
			j++
		} else {
			sum -= nums[i]
			i++
		}

		if sum >= s {
			minLen = min(minLen, j-i)
		}

	}
	return minLen % (n + 1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
