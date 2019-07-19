package problem1

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
//   为了避免
func twoSum(nums []int, target int) []int {

	hash := make(map[int]int, len(nums))

	for i, num := range nums {
		if _, ok := hash[target-num]; ok {
			return []int{hash[target-num], i}
		}
		hash[num] = i
	}
	return []int{}
}
