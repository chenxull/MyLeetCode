package problem448

/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 */
func findDisappearedNumbers(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	// 将出现的元素位置的下标都标记为负数
	for i := 0; i < len(nums); i++ {
		val := abs(nums[i]) - 1
		if nums[val] > 0 {
			// 数组中的值，可以当成是数组下标的索引
			nums[val] = -nums[val]
		}
	}

	res := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
