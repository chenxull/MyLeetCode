package problem75

/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */
//  使用三个指针进行操作，每个指针区间代表不同的颜色，根据出现的数字不同，来交换位置
func sortColors(nums []int) {
	i, j, k := 0, 0, len(nums)-1

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
