package offer39

func moreThanHalfNum(numbers []int) int {
	if numbers == nil {
		return 0
	}

	result := numbers[0]
	times := 1

	for i := 1; i < len(numbers); i++ {
		if times == 0 {
			result = numbers[i]
			times = 1
		} else if result == numbers[i] {
			times++
		} else {
			times--
		}
	}

	return result
}
