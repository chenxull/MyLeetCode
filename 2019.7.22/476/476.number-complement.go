package problem476

/*
 * @lc app=leetcode id=476 lang=golang
 *
 * [476] Number Complement
 */
//  异或^来求解，比如101 ^ 111 = 010。那么怎么得到111呢？考虑111 + 1 = 1000，而1000又是 最小的 大于101的 只有一位是1 的二进制数。
func findComplement(num int) int {
	// 使用异或的方式
	res := 0
	temp := num
	for temp > 0 {
		temp >>= 1
		//  构建 111
		res <<= 1
		res++
	}
	return num ^ res
}
