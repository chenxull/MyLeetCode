package problem401

import "fmt"

/*
 * @lc app=leetcode id=401 lang=golang
 *
 * [401] Binary Watch
 */
func readBinaryWatch(num int) []string {
	res := []string{}
	leds := make([]bool, 10)
	backtrack(num, 0, leds, &res)
	return res
}

func backtrack(num int, start int, leds []bool, res *[]string) {
	if num == 0 {
		m, h := get(leds[:6]), get(leds[6:])

		if m < 60 && h < 12 {
			*res = append(*res, fmt.Sprintf("%d:%02d", h, m))
		}
		return
	}

	for i := start; i < len(leds)-num+1; i++ {
		leds[i] = true
		backtrack(num-1, i+1, leds, res)
		leds[i] = false
	}

}

var bs = []int{1, 2, 4, 8, 16, 32}

func get(led []bool) int {
	sum := 0
	for i := 0; i < len(led); i++ {
		if led[i] {
			sum += bs[i]
		}
	}
	return sum
}
