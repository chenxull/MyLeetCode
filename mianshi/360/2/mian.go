package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	num := make([]int, m)
	for i := 0; i < m; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		num[i] = tmp
	}
	// fmt.Println(num)
	res := []int{}
	for i := 0; i < n; i++ {
		dfs(num, 0, i, n, i, &res)
	}
	fmt.Println(len(res))
}

func dfs(num []int, index, postion, n, step int, res *[]int) {
	if index >= len(num) && postion >= 0 && postion < n {
		*res = append(*res, postion)
		return
	}

	if num[index] <= postion {
		dfs(num, index+1, postion-num[index], n, step, res)
	}
	if num[index] <= n-postion-1 {
		dfs(num, index+1, postion+num[index], n, step, res)
	}
}
