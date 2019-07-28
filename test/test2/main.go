package main

import "fmt"

func main() {

	res := make([][]string, 4)
	// res = append(res, 1)
	res = append(res, []string{"sasd"})
	fmt.Println(res, len(res), cap(res))
}
