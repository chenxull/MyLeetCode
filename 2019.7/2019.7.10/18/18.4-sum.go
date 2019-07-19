package problem18

import "sort"

/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */
//  多了一层循环，转换为 3sum
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
					// 为了避免重复，l,r 都需要移动到不同的数字上
					l, r = removeDuplicate(nums, l, r)
				}
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
