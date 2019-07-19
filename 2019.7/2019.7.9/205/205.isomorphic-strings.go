package problem205

/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */
func isIsomorphic(s string, t string) bool {

	//  对应位置字母的数量相同，可以通过构建数组 hash 表
	if len(s) != len(t) {
		return false
	}

	m1 := make([]int, 256)
	m2 := make([]int, 256)

	for i := 0; i < len(s); i++ {
		if m1[int(s[i])] != m2[int(t[i])] {
			return false
		}
		//  关键之处，使用了 i+1 。这样就把数字与位置挂钩
		m1[int(s[i])] = i + 1
		m2[int(t[i])] = i + 1
	}
	return true
	// if len(s) != len(t) {
	// 	return false
	// }
	// bs := strings.Split(s, "")
	// bt := strings.Split(t, "")

	// if isPatter(bs, bt) && isPatter(bt, bs) {
	// 	return true
	// }
	// return false

}

// func isPatter(bs, bt []string) bool {
// 	size := len(bs)
// 	pattern := make(map[string]string, size)
// 	//  构建 map，如果构建成功
// 	for i := 0; i < size; i++ {
// 		if w, ok := pattern[bs[i]]; ok {
// 			if w != bt[i] {
// 				return false
// 			}
// 		} else {
// 			pattern[bs[i]] = bt[i]
// 		}
// 	}
// 	return true
// }
