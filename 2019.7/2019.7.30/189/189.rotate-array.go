package problem189

/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */
func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	size := len(nums)
	k = k % size
	rote(nums, size-k, size-1)
	rote(nums, 0, size-k-1)
	rote(nums, 0, size-1)
}

func rote(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
