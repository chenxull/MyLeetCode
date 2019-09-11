package problem415

/*
 * @lc app=leetcode id=415 lang=golang
 *
 * [415] Add Strings
 */
func addStrings(num1 string, num2 string) string {
	// num2 中保存着 更大的元素
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}

	len1 := len(num1)
	len2 := len(num2)
	// 将 string 转化为 byte
	byteNum1 := []byte(num1)
	byteNum2 := []byte(num2)

	carry := byte(0)
	ans := make([]byte, len2+1)
	// ans 最高位先置为 1
	ans[0] = '1'
	// 从最后一位开始进行处理
	i := 1
	for i <= len2 {
		if i <= len1 {
			ans[len2-i+1] = byteNum1[len1-i] - '0'
		}
		ans[len2-i+1] += byteNum2[len2-i] + carry

		if ans[len2-i+1] > '9' {
			ans[len2-i+1] -= 10
			carry = byte(1)
		} else {
			carry = byte(0)
		}
		i++
	}
	if carry == 1 {
		return string(ans)
	}
	return string(ans[1:])
}
