package main

import "fmt"

// 给定一个长度为R的数组，要求把其中的元素打散插入道长度为N的另外一个数组，
// 要求原先属于R的相邻元素之间的间隔和越大约好

func main() {
	numsR := []int{1, 2, 3}
	numsN := []int{4, 3, 5, 6, 8, 7, 9}
	sizeR := len(numsR)
	sizeN := len(numsN)
	length := sizeR + sizeN

	res := make([]int, 0, length)

	// 相邻元素间隔尽量相等
	part := sizeN / sizeR
	indexR := 0
	indexD := 1
	for i := 0; i < length; i++ {
		if i%part+1 == 0 {
			// 最后元素添加到res末尾
			if indexR == sizeR {
				res[length-1] = numsR[indexR]
			} else {
				res[i] = numsR[indexR]
				indexR++
			}
			continue
		}
		res[i] = numsN[indexD]
		indexD++
	}
	fmt.Println(res)
}
