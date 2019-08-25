package problem338

/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 */
func countBits(num int) []int {
	res := make([]int, num+1)
	res[0] = 0
	offset := 1
	// 注意找规律：2 4 8 16 中 1的个数都为 1 。其他数字的二进制中 1 的个数，都可以组合而成
	for i := 1; i < num+1; i++ {
		if offset*2 == i {
			offset *= 2
		}
		res[i] = res[i-offset] + 1
	}
	return res

}
