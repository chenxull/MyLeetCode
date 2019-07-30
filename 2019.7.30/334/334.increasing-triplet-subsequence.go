package problem334

/*
 * @lc app=leetcode id=334 lang=golang
 *
 * [334] Increasing Triplet Subsequence
 */
func increasingTriplet(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	max := 1<<63 - 1
	small, big := max, max
	for _, n := range nums {
		if n <= small {
			small = n
		} else if n < big {
			big = n
		} else {
			return true
		}
	}

	return false
}
