package problem739

/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */
// 还需要重新研究
func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		result := check(T[i], i, T)
		res[i] = result
	}
	return res

}

func check(input, i int, T []int) int {
	for j := i; j < len(T); j++ {
		if T[j] > input {
			return j - i
		} else {
			continue
		}
	}
	return 0
}
