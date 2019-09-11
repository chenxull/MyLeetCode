package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		nums = append(nums, tmp)
	}
	fmt.
}
