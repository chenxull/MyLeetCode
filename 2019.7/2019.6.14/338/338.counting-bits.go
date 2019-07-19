package problem338

/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 */
func countBits(num int) []int {
	res := make([]int, num+1)

	for i := 1; i <= num; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

//res := make([]int, num+1)
//res[0] = 0
//offset := 1
//for i := 1; i < num+1; i++ {
//	if offset*2 == i {
//		offset = offset * 2
//	}
//	res[i] = res[i-offset] + 1
//}
//
//return res
