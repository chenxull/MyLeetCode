package problem594

/*
 * @lc app=leetcode id=594 lang=golang
 *
 * [594] Longest Harmonious Subsequence
 */
func findLHS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	hashmap := make(map[int]int, len(nums))

	for _, v := range nums {
		hashmap[v]++
	}

	m := 0

	for k, v := range hashmap {
		v1, ok := hashmap[k+1]

		if ok {
			if v+v1 > m {
				m = v + v1
			}
		}
	}
	return m

}
