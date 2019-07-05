package problem38

/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */
func countAndSay(n int) string {

	buf := []byte{'1'}

	for n > 1 {
		buf = say(buf)
		n--
	}

	return string(buf)

}

func say(buf []byte) []byte {
	res := make([]byte, 0, 2*len(buf))
	i, j := 0, 1
	for i < len(buf) {
		for j < len(buf) && buf[i] == buf[j] {
			j++
		}
		// res = append(res, byte(j-i+'0'))
		// res = append(res, buf[i])
		// // fmt.Println(res)
		res = append(res, byte(j-i+'0'), buf[i])
		i = j
	}
	return res
}
