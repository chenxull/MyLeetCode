package problem18

import "sort"

/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */
//  和 3sum 类似，首先培训，然后固定二个元素，对剩下的二个元素进行双指针对撞求解
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			l, r := j+1, len(nums)-1

			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				switch {
				case sum < target:
					l++
				case sum > target:
					r--
				default:
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l, r = deduplicate(nums, l, r)
				}
			}
		}
	}
	return res
}

func deduplicate(nums []int, i, j int) (int, int) {
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
