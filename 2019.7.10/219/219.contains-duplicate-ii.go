package problem219

/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */
func containsNearbyDuplicate(nums []int, k int) bool {
	hash := make(map[int]int, len(nums))
	for i, n := range nums {
		if hash[n] != 0 && i-(hash[n]-1) <= k {
			return true
		}
		// 如果 hash[n] =i ，假如重复的是第一个元素，就会无法判断
		hash[n] = 1

	}
	return false
}
