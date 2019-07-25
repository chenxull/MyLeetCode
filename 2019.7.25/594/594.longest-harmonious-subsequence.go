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

	longest := 0

	//  可以将 这个 for 循环换成，遍历 hashmap，可以大大提高效率
	// for i := 1; i < len(nums); i++ {
	// 	currnetNum := nums[i]
	// 	currentLen := hashmap[currnetNum]
	// 	if hashmap[currnetNum-1] != 0 && currnetNum != nums[i-1] {
	// 		currentLen += hashmap[currnetNum-1]
	// 		longest = max(longest, currentLen)

	// 	}
	// }

	for k, v1 := range hashmap {
		v2, ok := hashmap[k+1]
		if ok {
			t := v1 + v2
			if longest < t {
				longest = t
			}
		}
	}
	return longest
}
