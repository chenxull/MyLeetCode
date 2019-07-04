package main

import "fmt"

func main() {
	s := "abcdefg"
	for a, b := range s {
		fmt.Println("a:", a, "b:", b)
	}
}
