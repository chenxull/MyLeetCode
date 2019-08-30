package problem219

/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */
func containsNearbyDuplicate(nums []int, k int) bool {
	if k < 0 {
		return false
	}
	hashmap := make(map[int]int, len(nums))

	for i, n := range nums {
		//  第二次遇到
		if hashmap[n] != 0 && i-(hashmap[n]-1) <= k {
			return true
		}
		hashmap[n] = 1 + i
	}
	return false
}
