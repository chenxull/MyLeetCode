package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 2, 56, 67}

	for i := range s {
		fmt.Println(i)
	}
}
