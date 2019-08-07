package problem11

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
//  使用双指针对撞的思想，使用一个变量保存出现的最大面积

func maxArea(height []int) int {
	lo, hi := 0, len(height)-1
	maxare := 0
	for lo < hi {
		weight := hi - lo
		h := min(height[lo], height[hi])
		area := h * weight
		if area > maxare {
			maxare = area
		}

		// 移动双指针，当其中一遍的指针偏小时，启动这边的指针，寻求更大的值
		if height[lo] < height[hi] {
			lo++
		} else {
			hi--
		}
	}
	return maxare
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
