package problem1

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	hashmap := make(map[int]int, len(nums))

	// 为什么这种赋值方式会导致几个测试用例无法通过
	// for i, b := range nums {
	// 	hashmap[b] = i
	// }

	for idx, v := range nums {
		if i, ok := hashmap[target-v]; ok {
			return []int{i, idx}
		}
		hashmap[v] = idx

	}
	return nil
}
