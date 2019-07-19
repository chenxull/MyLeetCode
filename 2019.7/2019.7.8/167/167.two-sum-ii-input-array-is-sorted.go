package problem167

/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input array is sorted
 */
func twoSum(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return []int{}
	}
	i, j := 0, len(numbers)-1

	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}

		if sum < target {
			i++
		} else {
			j--
		}
	}

	return []int{}

}
