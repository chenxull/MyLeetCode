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

	index, i, count, end := 1, 0, 0, len(nums)-1

	var nextI, maxNextI, maxI int

	for index < len(nums)/2 {
		i = index
		for i < end {
			if i+nums[i] >= end {
				fmt.Println(count + 1)
				return
			}

			nextI, maxNextI = i+1, i+nums[i]
			for nextI <= maxNextI && nextI < end {

				if nextI+nums[nextI] > maxI {
					maxI, i = nextI+nums[nextI], nextI
				}

				nextI++
			}

			count++
		}
		index++
	}

	if count == 0 {
		fmt.Println(-1)
		return
	}
	// fmt.Println(count)

}
