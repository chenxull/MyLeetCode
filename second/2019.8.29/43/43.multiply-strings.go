package problem43

/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */
func multiply(num1 string, num2 string) string {

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	n1 := []byte(num1)
	n2 := []byte(num2)

	temp := make([]int, len(n1)+len(n2))

	for i := 0; i < len(n1); i++ {
		for j := 0; j < len(n2); j++ {
			temp[i+j+1] += int(n1[i]-'0') * int(n2[j]-'0')
		}
	}

	for i := len(temp) - 1; i > 0; i-- {
		temp[i-1] += temp[i] / 10
		temp[i] = temp[i] % 10
	}

	if temp[0] == 0 {
		temp = temp[1:]
	}

	//  将 int 转换为 []byte
	res := make([]byte, len(temp))

	for i := 0; i < len(res); i++ {
		res[i] = byte(temp[i]) + '0'
	}
	return string(res)
}
