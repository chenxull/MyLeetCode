package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	nums1 := make([]int, n)
	nums2 := make([]int, m)

	for i := 0; i < n; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		nums1[i] = tmp
	}

	for i := 0; i < m; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		nums2[i] = tmp
	}
	// fmt.Println(nums1, nums2)
	oddCountN := 0
	evenCountN := 0
	for i := 0; i < n; i++ {
		if nums1[i]%2 == 1 {
			oddCountN++
		} else {
			evenCountN++
		}
	}
	oddCountM := 0
	evenCountM := 0
	for i := 0; i < m; i++ {
		if nums2[i]%2 == 1 {
			oddCountM++
		} else {
			evenCountM++
		}
	}

	fmt.Println(min(oddCountN, evenCountM) + min(oddCountM, evenCountN))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
