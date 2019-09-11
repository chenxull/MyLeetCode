package main

import "fmt"

//请你在s的所有子串中，找到出现次数最多的子串，告诉老师它的出现次数。
func main() {
	var str string
	fmt.Scan(&str)
	// fmt.Println(str)
	hashmap := make(map[string]int)

	res := []string{}
	cur := ""
	for i := 1; i <= len(str); i++ {
		backtrack(str, cur, 0, i, &res)
	}

	for _, v := range res {
		hashmap[v]++
	}
	max := 0
	for _, count := range hashmap {
		if count > max {
			max = count
		}
	}
	fmt.Println(max)
	// fmt.Println(res)
}

func backtrack(str, cur string, start, n int, res *[]string) {
	if len(cur) == n {
		*res = append(*res, cur)
		return
	}

	for i := start; i < len(str); i++ {
		if i+n > len(str) {
			return
		}
		cur += str[i : i+n]
		backtrack(str, cur, i+1, n, res)
		cur = ""
	}
}
