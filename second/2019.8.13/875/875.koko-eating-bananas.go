package problem875

/*
 * @lc app=leetcode id=875 lang=golang
 *
 * [875] Koko Eating Bananas
 */
func minEatingSpeed(piles []int, H int) int {
	l := 1
	r := findMax(piles)

	for l < r {
		m := l + (r-l)/2
		// canEatall 就是 g(m)函数
		if canEatall(piles, m, H) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func canEatall(piles []int, m, H int) bool {
	// h 以当前速度吃香蕉所需要的时间
	h := 0
	for _, p := range piles {
		// 这里有点绕
		h += (p + m - 1) / m
	}
	if h > H {
		return false
	}
	return true
}

func findMax(nums []int) int {
	mx := 0
	for _, n := range nums {
		mx = max(n, mx)
	}
	return mx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
