package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([][]int, n)
	for i := range nums {
		nums[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var num int
			fmt.Scan(&num)
			// nums[i] = append(nums[i], num)
			nums[i][j] = num
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if nums[i][j] > 3 {

			}
		}
	}

}
