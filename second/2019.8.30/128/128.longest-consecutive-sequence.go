package problem

/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */
func longestConsecutive(nums []int) int {
	// 使用空间换时间
	hashmap := make(map[int]bool, len(nums))

	for i := range nums {
		hashmap[nums[i]] = true
	}
	longest := 0
	for i := 0; i < len(nums); i++ {
		// 如果存在连续数字，从最小的开始计算
		if !hashmap[nums[i]-1] {
			currentNum := nums[i]
			currentLen := 1
			for hashmap[currentNum+1] {
				currentLen++
				currentNum++
			}
			longest = max(longest, currentLen)
		}
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
