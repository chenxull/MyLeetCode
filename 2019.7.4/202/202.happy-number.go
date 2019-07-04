package problem202

/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */
func isHappy(n int) bool {
	slow, fast := n, trans(n)
	//  slow fast 有二种可能退出循环， 都为 1 或进入循环
	for slow != fast {
		slow = trans(slow)
		fast = trans(trans(fast))
	}

	if slow == 1 {
		return true
	}

	return false
}

//  求数字的和
func trans(n int) int {
	res := 0
	for n != 0 {
		res += (n % 10) * (n % 10)
		n = n / 10
	}
	return res
}
