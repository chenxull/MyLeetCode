package main

import "fmt"

/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */
//  有点复杂
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	//  l 中保存长一点的长度
	l := len(a)

	//  确保创建出来的数组，有相同的长度
	isA := trans(a, l)
	isB := trans(b, l)

	return makeString(add(isA, isB))
}

// 将 string 转换为 int
func trans(s string, length int) []int {
	res := make([]int, length)
	ls := len(s)
	// 构建新的参数数组,较小的那个数组，高位会用 0 补齐
	for i, v := range s {
		res[length-ls+i] = int(v - '0')
	}
	return res
}

// 相加
func add(a, b []int) []int {
	// l 的长度考虑到了进位
	l := len(a) + 1
	res := make([]int, l)
	// 因为数组的高位保存的是较小的数字,从高位开始计算
	for i := l - 1; i >= 1; i-- {
		temp := res[i] + a[i-1] + b[i-1]
		res[i] = temp % 2   // 对进位的处理，如果不进位保持原样，如果进位为 0
		res[i-1] = temp / 2 //对高一位的处理，如果 temp =2 高一位进位成1，反之为 0

	}

	// 对高位 0 的去除处理
	i := 0
	for i < l-1 && res[i] == 0 {
		i++
	}

	return res[i:]
}

//  int 装换为 string
func makeString(nums []int) string {
	res := make([]byte, len(nums))
	for i := range nums {
		res[i] = byte(nums[i] + '0')
	}
	return string(res)
}

func main() {
	a := "1010"
	b := "0011"
	fmt.Println(addBinary(a, b))
}
