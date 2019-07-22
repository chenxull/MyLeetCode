package problem762

/*
 * @lc app=leetcode id=762 lang=golang
 *
 * [762] Prime Number of Set Bits in Binary Representation
 */
func countPrimeSetBits(L int, R int) int {
	//  使用一个数组保存质数信息,没有标记到的，都为 0
	primes := [...]int{2: 1, 3: 1, 5: 1, 7: 1, 11: 1, 13: 1, 17: 1, 19: 1}

	res := 0
	for i := L; i <= R; i++ {
		bits := 0
		// 统计每个数字的二进制中有多少个 1
		for n := i; n > 0; n >>= 1 {
			bits += n & 1
		}
		res += primes[bits]
	}
	return res
}
