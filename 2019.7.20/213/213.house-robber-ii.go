package problem213

/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 */
func rob(nums []int) int {
	n := len(nums)

	switch n {
	case 0:
		return 0
	case 1:
		return nums[0]
	}

	// 这题的关键在于圆环的切割
	return max(robbing(nums[1:]), robbing(nums[:n-1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func robbing(nums []int) int {
	var a, b int
	for i, num := range nums {
		if i%2 == 0 {
			a = max(a+num, b)
		} else {
			b = max(b+num, a)
		}
	}
	return max(a, b)
}
