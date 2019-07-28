package main

import "fmt"

func main() {
	res := make(map[int]int, 5)
	res[3] = 1
	res[2] = 4

	for i, n := range res {
		fmt.Println(i, n)
	}
}
