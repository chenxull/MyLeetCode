package problem258


/*
 * @lc app=leetcode id=258 lang=golang
 *
 * [258] Add Digits
 */
func addDigits(num int) int {
	if num <= 9 {
		return num
	}

	for num > 9 {
		num = add(num)
	}

	return num
}

func add(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

