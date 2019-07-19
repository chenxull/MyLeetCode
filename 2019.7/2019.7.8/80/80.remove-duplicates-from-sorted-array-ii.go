package problem80

/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0

	for j := 0; j < len(nums); j++ {
		//  前二次直接 pass
		if i < 2 || nums[j] != nums[i-2] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

// 可以把 a 想象成一层一层摆好的木板，从下往上，索引值变大
// 重复的多余木板被抽掉，上面的木板就会下落
//
// i 始终指向下一个合格木板要放置的位置
// j 始终指向下一个需要检查的木板的位置
//
// 注意前提条件： a is sorted
//
// a[i-2] != a[j] 说明 a[j] 是合格的木板，
// a[j] 重复次数还没有超过 2 次
//
// a[i] = a[j] 就是木板下落
//
// 到最后，i 就成了所有合格木板的数量，也就是 a 合格部分的长度
