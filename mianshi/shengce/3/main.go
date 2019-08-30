package main

import "fmt"

func main() {
	var n, m int
	a := make([]int, n)
	b := make([]int, m)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		a[i] = tmp
	}

	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		b[i] = tmp
	}

	fmt.Println(a, b)
}
