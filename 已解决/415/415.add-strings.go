package problem415

/*
 * @lc app=leetcode id=415 lang=golang
 *
 * [415] Add Strings
 *
 * https://leetcode.com/problems/add-strings/description/
 *
 * algorithms
 * Easy (42.96%)
 * Total Accepted:    90.5K
 * Total Submissions: 208.3K
 * Testcase Example:  '"0"\n"0"'
 *
 * Given two non-negative integers num1 and num2 represented as string, return
 * the sum of num1 and num2.
 *
 * Note:
 *
 * The length of both num1 and num2 is < 5100.
 * Both num1 and num2 contains only digits 0-9.
 * Both num1 and num2 does not contain any leading zero.
 * You must not use any built-in BigInteger library or convert the inputs to
 * integer directly.
 *
 *
 */
func addStrings(num1 string, num2 string) string {
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}

	len1, len2 := len(num1), len(num2)
	byte1, byte2 := []byte(num1), []byte(num2)

	carry := byte(0)
	ans := make([]byte, len2+1)
	ans[0] = '1'
	i := 1
	for i <= len2 {
		if i <= len1 {
			ans[len2+1-i] = byte1[len1-i] - '0'
		}
		ans[len2+1-i] += byte2[len2-i] + carry

		if ans[len2+1-i] > '9' {
			ans[len2+1-i] -= 10
			carry = byte(1)
		} else {
			carry = byte(0)
		}
		//fmt.Println("ans的长度为", len(ans))
		i++
	}
	if carry == 1 {
		return string(ans)
	}
	return string(ans[1:])
}
