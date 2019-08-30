package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)

	if x == 1 {
		fmt.Println(2)
		return
	} else if x == -1 {
		fmt.Println(4)
		return
	} else if x == 0 {
		fmt.Println(-1)
		return

	}
	fmt.Println(-1)

}
