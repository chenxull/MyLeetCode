// package problem788

/*
 * @lc app=leetcode id=788 lang=golang
 *
 * [788] Rotated Digits
 */
func rotatedDigits(N int) int {
	res := 0

	for i := 2; i <= N; i++ {
		flag := isValid(i)
		if flag {
			res++
		}
	}
	return res
}

func isValid(a int) bool {
	flag := false

	for a > 0 {
		switch a % 10 {
		case 2, 5, 6, 9:
			flag = true
		case 3, 4, 7:
			return false
		}
		a /= 10
	}
	return flag
}
