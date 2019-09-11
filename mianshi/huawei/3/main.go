package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	strs := scanner.Text()
	strnums := strings.Split(strs, " ")



	



	nums := make([]int, 0, len(strnums))
	for _, v := range strnums {
		num, _ := strconv.Atoi(v)
		nums = append(nums, num)

	}

	mincount, count, end := 100, 0, len(nums)-1
	var nextI, maxNextI, maxI int

	for i := 1; i < len(nums)/2; i++ {
		for i < end {
			if i+nums[i] >= end {
				mincount = min(mincount, count+1)
				break
			}

			nextI, maxNextI = i+1, i+nums[i]
			for nextI <= maxNextI {

				if nextI+nums[nextI] > maxI {
					maxI, i = nextI+nums[nextI], nextI
				}

				nextI++
			}

			count++
		}
	}

	if mincount == 0 {
		fmt.Println(-1)
		return
	}
	fmt.Println(mincount)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
