package offer44

import "math"

func countOfIntegers(digit int) int {
	if digit == 1 {
		return 10
	}

	return int(math.Pow10(digit-1) * 9)
}

func beginNumber(digit int) int {
	if digit == 1 {
		return 0
	}
	return int(math.Pow10(digit - 1))
}

func digitAtIndexCore(index, digit int) int {
	// 对给定元素镜像操作
	number := beginNumber(digit) + index/digit
	indexFromRight := digit - index%digit
	for i := 1; i < indexFromRight; i++ {
		number /= 10
	}
	return number % 10
}

func digitAtIndex(index int) int {
	if index < 0 {
		return 0
	}
	digit := 1
	for {
		// 计算位数 例如三位数是 100
		numbers := countOfIntegers(digit)

		// 当要找的那一位数字位于某 m 位数之中后，使用如下函数找出具体的位置。
		if index < numbers*digit {
			return digitAtIndexCore(index, digit)
		}
		index -= numbers * digit
		digit++
	}
	return -1
}
