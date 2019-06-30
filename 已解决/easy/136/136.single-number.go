package problem136

/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */
func singleNumber(nums []int) int {
	//ans := 0
	//sort.Ints(nums)
	//for i := 0; i < len(nums); i += 2 {
	//	if i == len(nums)-1 {
	//		return nums[i]
	//	}
	//	if nums[i] == nums[i+1] {
	//		continue
	//	}
	//	ans = nums[i]
	//	break
	//}
	//return ans
	res := 0
	for _, n := range nums {
		res ^= n
	}
	return res
}
