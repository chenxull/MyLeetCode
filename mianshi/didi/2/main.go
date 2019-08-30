package main

import "fmt"

func main() {
	var n, cost, total int
	fmt.Scan(&n, &total, &cost)

	nums := make([][]int, 2)

	for i := 0; i < 2; i++ {
		nums[i] = make([]int, n)
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < n; j++ {
			temp := 0
			fmt.Scan(&temp)
			nums[i][j] = temp
		}
	}
	fmt.Println(nums)
}
