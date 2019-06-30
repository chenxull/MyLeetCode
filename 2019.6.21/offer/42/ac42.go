package offer42

func FindGreatestSumOfSubArray(data []int) int {
	if data == nil {
		return -1
	}

	var currentSum, greatestSum int
	// 一个循环分为二个阶段。 一是处理当前值的和，而是当前值的和与目前保存的最大值之间的比较
	for i := 0; i < len(data); i++ {
		if currentSum <= 0 {
			currentSum = data[i]
		} else {
			currentSum += data[i]
		}

		if currentSum > greatestSum {
			greatestSum = currentSum
		}
	}
	return greatestSum
}
