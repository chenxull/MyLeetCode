package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	num := make([]int, 0, n)
	for i := 0; i < n; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		num = append(num, tmp)
	}

	res := [][]int{}
	token := make([]bool, n)
	cur := make([]int, n)
	backtrack(num, 0, cur, token, &res)
	// fmt.Println(res)

	sort.Slice(res, func(i, j int) bool {
		num1 := res[i]
		num2 := res[j]
		str1 := ""
		str2 := ""
		for i := 0; i < len(num1); i++ {
			str1 += strconv.Itoa(num1[i])
			str2 += strconv.Itoa(num2[i])

		}

		num11, _ := strconv.Atoi(str1)
		num22, _ := strconv.Atoi(str2)

		return num11 < num22
	})

	for i := 0; i < len(res); i++ {
		ans := ""
		for j := 0; j < len(res[0]); j++ {
			ans += strconv.Itoa(res[i][j])
		}
		fmt.Print(ans)
		if i == len(res)-1 {
			break
		} else {
			fmt.Print(" ")
		}
	}
}

func backtrack(num []int, index int, cur []int, token []bool, res *[][]int) {
	if index == len(num) {
		tmp := make([]int, len(num))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(num); i++ {
		if !token[i] {
			token[i] = true
			cur[index] = num[i]
			backtrack(num, index+1, cur, token, res)
			token[i] = false
		}
	}
}
