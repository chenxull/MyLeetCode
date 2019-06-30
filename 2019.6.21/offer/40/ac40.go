package offer40

import "sort"

func getLeastNumbers(data []int, k int) []int {
	LeastNumbers := make([]int, k)
	for i := 0; i < k; i++ {
		LeastNumbers[i] = data[i]
	}

	sort.Ints(LeastNumbers)

	for i := k; i < len(data); i++ {
		if data[i] < LeastNumbers[k-1] {
			LeastNumbers[k-1] = data[i]
			sort.Ints(LeastNumbers)
		}
	}
	return LeastNumbers
}
