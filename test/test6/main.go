package main

import "fmt"

func dfs(gird [][]int) {
	for i := 0; i < len(gird); i++ {
		for j := 0; j < len(gird[0]); j++ {
			gird[i][j] = 0
		}
	}
}

func main() {
	gird := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	dfs(gird)

	fmt.Println(gird)
}
