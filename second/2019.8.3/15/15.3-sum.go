package problem15

import "sort"

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
//  首先对原有的数组进行排序，然后固定一个元素，对剩下的二个元素使用
// 双指针的思路，寻找到服务要求的解。同时在获取到符合到要去的解之后，
// pass 掉相同元素
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

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
				l, r = deduplicate(nums, l, r)
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
