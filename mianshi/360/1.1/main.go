package main

import "fmt"

//请你在s的所有子串中，找到出现次数最多的子串，告诉老师它的出现次数。
func main() {
	var str string
	fmt.Scan(&str)
	// fmt.Println(str)
	hashmap := make(map[byte]int)

	// res := []string{}

	for i := 0; i < len(str); i++ {
		hashmap[str[i]]++
	}

	max := 0
	for _, count := range hashmap {
		if count > max {
			max = count
		}
	}
	fmt.Println(max)
}
