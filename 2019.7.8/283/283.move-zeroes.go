// package problem283

/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */
func moveZeroes(nums []int) {
	if nums == nil {
		return
	}

	lastNonZeroFoundAt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			swap(i, lastNonZeroFoundAt, nums)
			lastNonZeroFoundAt++
		}
	}

}

func swap(a, b int, nums []int) {
	nums[a], nums[b] = nums[b], nums[a]
}
