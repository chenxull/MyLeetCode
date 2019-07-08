package problem75

/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */
func sortColors(nums []int) {
	i, j, k := 0, 0, len(nums)-1

	//  [0,i] =0   [i,j] =1 [j,k] = 2
	for j <= k {
		switch nums[j] {
		case 0:
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		case 1:
			j++
		case 2:
			nums[j], nums[k] = nums[k], nums[j]
			k--
		}
	}
}
