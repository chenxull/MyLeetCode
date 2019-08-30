package problem150

import "strconv"

/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */
func evalRPN(tokens []string) int {
	nums := make([]int, 0, len(tokens))
	for _, i := range tokens {
		if i == "+" || i == "/" || i == "-" || i == "*" {
			a, b := nums[len(nums)-1], nums[len(nums)-2]
			nums = nums[:len(nums)-2]
			nums = append(nums, cal(a, b, i))
		} else {
			num, _ := strconv.Atoi(i)
			nums = append(nums, num)
		}
	}
	return nums[0]
}

func cal(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "/":
		return b / a
	case "*":
		return a * b
	default:
		return b - a
	}

}
