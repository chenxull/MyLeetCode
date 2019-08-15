package problem80

/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 新数组中最后一位
	index := 0

	// 遍历整个数组，在判断的同时处理移动整个数组
	for i := 0; i < len(nums); i++ {
		// a.最开始的二个元素直接移动位置
		// b.之后每次移动之前，需要判断当前位置元素 i 是否在新数组中出现次数操作 2 次

		if index < 2 || nums[i] != nums[index-2] {
			nums[index] = nums[i]
			index++ //移动新数组中最后一位的下标
		}
		// 没有进入上述 if 语句，说明有元素出现的次数超过了 2 次。
	}
	return index
}
