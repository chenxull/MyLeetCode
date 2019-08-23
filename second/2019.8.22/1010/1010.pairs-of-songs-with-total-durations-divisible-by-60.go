package problem60

/*
 * @lc app=leetcode id=1010 lang=golang
 *
 * [1010] Pairs of Songs With Total Durations Divisible by 60
 */
//  一种新的处理方法
func numPairsDivisibleBy60(time []int) int {
	rec := [60]int{}
	for _, v := range time {
		rec[v%60]++
	}
	res := rec[0] * (rec[0] - 1) / 2

	res += rec[30] * (rec[30] - 1) / 2
	for i := 1; i < 30; i++ {
		res += rec[i] * rec[60-i]
	}
	return res
}
