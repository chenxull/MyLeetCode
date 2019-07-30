package problem594

/*
 * @lc app=leetcode id=594 lang=golang
 *
 * [594] Longest Harmonious Subsequence
 */
func findLHS1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	hashmap := make(map[int]int, len(nums))

	for _, v := range nums {
		hashmap[v]++
	}
	max := 0
	for k, v1 := range hashmap {
		// 学习整个写法
		v2, ok := hashmap[k+1]

		if ok {
			t := v1 + v2
			if max < t {
				max = t
			}
		}
	}
	return max
}
