package problem11

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
func maxArea(height []int) int {
	lo, hi := 0, len(height)-1
	maxarea := 0
	for lo <= hi {

		weight := hi - lo
		h := min(height[lo], height[hi])
		area := weight * h

		if area > maxarea {
			maxarea = area
		}

		// 变化趋势
		if height[hi] > height[lo] {
			lo++
		} else {
			hi--
		}
	}
	return maxarea
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
