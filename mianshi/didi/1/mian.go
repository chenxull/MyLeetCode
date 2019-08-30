package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int

	fmt.Scan(&n)
	size := 2*n - 1
	numsInput := make([]string, size+5)
	// 获取输入
	for i := 0; i < size; i++ {
		temp := ""
		fmt.Scan(&temp)
		numsInput[i] = temp
		// numsInput = append(numsInput, temp)
	}
	countA := 0
	countB := 0
	for i := 0; i < size-1; i += 2 {
		if bigger(numsInput[i], numsInput[i+2]) && numsInput[i+1] == "+" && numsInput[i+3] != "*" && numsInput[i+3] != "-" && numsInput[i+3] != "/" {
			numsInput[i], numsInput[i+2] = numsInput[i+2], numsInput[i]
			countA++
		} else if bigger(numsInput[i], numsInput[i+2]) && numsInput[i+1] == "*" && numsInput[i+3] != "/" {
			numsInput[i], numsInput[i+2] = numsInput[i+2], numsInput[i]
			countB++
		}
	}

	if countA < countB {
		countA = countB
	}

	for countA > 0 {
		for i := 0; i < size-1; i += 2 {
			if bigger(numsInput[i], numsInput[i+2]) && numsInput[i+1] == "+" && numsInput[i+3] != "*" && numsInput[i+3] != "-" && numsInput[i+3] != "/" {
				numsInput[i], numsInput[i+2] = numsInput[i+2], numsInput[i]

			} else if bigger(numsInput[i], numsInput[i+2]) && numsInput[i+1] == "*" && numsInput[i+3] != "/" {
				numsInput[i], numsInput[i+2] = numsInput[i+2], numsInput[i]
			}
		}
		countA--
	}

	fmt.Println(strings.Join(numsInput[:size], " "))

}

func bigger(a, b string) bool {

	inta, _ := strconv.Atoi(a)
	intb, _ := strconv.Atoi(b)

	if inta > intb {
		return true
	}
	return false
}
