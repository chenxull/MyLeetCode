package problem560

/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 */
func subarraySum(nums []int, k int) int {

	count, sum := 0, 0
	rec := make(map[int]int, len(nums))
	rec[0] = 1
	for i := range nums {
		sum += nums[i]
		// sum+nums[i]-k = sum
		count += rec[sum-k]
		rec[sum]++
	}
	return count

	// 普通遍历
	// count := 0
	// for i := 0; i < len(nums); i++ {
	// 	sum := 0
	// 	for j := i; j < len(nums); j++ {
	// 		sum += nums[j]
	// 		if sum == k {
	// 			count++
	// 		}
	// 	}
	// }
	// return count
}
