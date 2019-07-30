package problem693

/*
 * @lc app=leetcode id=693 lang=golang
 *
 * [693] Binary Number with Alternating Bits
 */
func hasAlternatingBits(n int) bool {
	// 因为 只有二种模式 ，10 或 01 模式，所以在一开始确定好后面 0，1 出来的模式
	std := n & 3 // 3 =11

	// 如果不为二种模式中的任一个，返回 false
	if std != 1 && std != 2 {
		return false
	}

	for n > 0 {
		if n&3 != std {
			return false
		}
		//  每次移动二位
		n >>= 2
	}
	return true
}
