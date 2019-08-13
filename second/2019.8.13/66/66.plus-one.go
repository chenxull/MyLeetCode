package problem66

/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */
//  关键在于处理进位以及首位进位
func plusOne(digits []int) []int {
	length := len(digits)

	if length == 0 {
		return nil
	}
	// 加 1
	digits[length-1]++

	// 处理进位
	for i := length - 1; i > 0; i-- {
		if digits[i] < 10 {
			break
		}
		digits[i] -= 10
		digits[i-1]++
	}

	// 处理首位
	if digits[0] > 9 {
		digits[0] -= 10
		digits = append([]int{1}, digits...)
	}
	return digits
}
