package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	//str:
	str1 := make([]string, 0, N)
	for i := 0; i < N; i++ {

		str1 = append(str1)
	}

	fmt.Println("CZ7132,A2,ZHAOSI")
	fmt.Println("CZ7156,A2,ZHANGSAN")
	fmt.Println("CZ7156,A3,WANGWU")
}
