package main

import (
	"LeetCode/sort"
	"fmt"
	"time"
)

func merge(left, right []int) []int {
	m, n := len(left), len(right)
	res := make([]int, m+n)

	j, k := 0, 0
	for i := 0; i < m+n; i++ {
		if k >= n {
			res[i] = left[j]
			j++
			continue
		}

		if j >= m {
			res[i] = right[k]
			k++
			continue
		}
		if left[j] < right[k] {
			res[i] = left[j]
			j++
		} else {
			res[i] = right[k]
			k++
		}
	}
	return res

	// for len(left) > 0 || len(right) > 0 {
	// 	if len(left) == 0 {
	// 		return append(res, right...)
	// 	}

	// 	if len(right) == 0 {
	// 		return append(res, left...)
	// 	}

	// 	if left[0] < right[0] {
	// 		res = append(res, left[0])
	// 		left = left[1:]
	// 	} else {
	// 		res = append(res, right[0])
	// 		right = right[1:]
	// 	}
	// }
	// return res
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return merge(left, right)
}

func main() {
	nums := utils.ArrayRand(100000)
	t1 := time.Now()
	fmt.Println(mergeSort(nums))
	during := time.Since(t1)
	fmt.Println(during)
}
