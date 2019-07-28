package main

import "fmt"

func main() {
	s := "12345"
	b := []byte(s)
	b[1], b[2] = b[3], b[4]
	fmt.Println(string(b))
}
