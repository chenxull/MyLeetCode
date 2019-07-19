package problem150

import "strconv"

/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */
func evalRPN(tokens []string) int {
	nums := make([]int, 0, len(tokens))

	for _, s := range tokens {
		if s == "+" || s == "-" || s == "/" || s == "*" {
			b, a := nums[len(nums)-1], nums[len(nums)-2]

			nums = nums[:len(nums)-2]
			operator := s
			//  将计算过的元素放入 stack 中
			nums = append(nums, calculate(operator, a, b))
		} else {
			num, _ := strconv.Atoi(s)
			nums = append(nums, num)
		}
	}
	return nums[0]
}

func calculate(operator string, a, b int) int {
	ans := 0
	switch operator {
	case "+":
		ans = a + b
	case "-":
		ans = a - b
	case "*":
		ans = a * b
	default:
		ans = a / b
	}
	return ans
}
