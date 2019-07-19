package problem15

import "sort"

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
//   这题的关键在于去重复
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := range nums {
		//  避免 i 所在的元素出现重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 因为剩下的数组是有序的，可以使用双指针对撞的方法求解。
		l, r := i+1, len(nums)-1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			switch {
			case sum < 0:
				l++
			case sum > 0:
				r--
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 为了避免重复，l,r 都需要移动到不同的数字上
				l, r = removeDuplicate(nums, l, r)
			}
		}
	}
	return res
}

func removeDuplicate(nums []int, i, j int) (int, int) {
	for i < j {
		switch {
		case nums[i] == nums[i+1]:
			i++
		case nums[j] == nums[j-1]:
			j--
		default:
			i++
			j--
			return i, j
		}
	}
	return i, j
}
