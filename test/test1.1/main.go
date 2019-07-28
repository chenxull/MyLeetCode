package main

import "fmt"

func main() {
	res := 0
	temp := 5
	for temp > 0 {
		temp >>= 1
		res <<= 1
		res++
	}
	fmt.Println(res)
}
