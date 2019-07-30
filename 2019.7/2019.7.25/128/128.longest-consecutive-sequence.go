package problem128

/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */
//   构建一个 map 来存储出现的元素
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	hashmap := make(map[int]bool, len(nums))
	for _, v := range nums {
		hashmap[v] = true
	}

	longest := 0
	for i := 0; i < len(nums); i++ {
		if !hashmap[nums[i]-1] {
			currentNum := nums[i]
			currentLen := 1
			for hashmap[currentNum+1] {
				currentNum++
				currentLen++
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
