package problem43

// package main

// import "fmt"

/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	n1 := []byte(num1)
	n2 := []byte(num2)

	// 使用[]int 类型，是因为 byte 类型最大值为255，可能会越界
	temp := make([]int, len(n1)+len(n2))

	//  从高位开始计算，对应位的结果保存在对应的temp 数组中
	for i := 0; i < len(n1); i++ {
		for j := 0; j < len(n2); j++ {
			//  +1 是给最高位留一个进位空间。 temp中存储的都是对应位的乘积结果
			temp[i+j+1] += int(n1[i]-'0') * int(n2[j]-'0')
		}
	}

	// 处理进位
	for i := len(temp) - 1; i > 0; i-- {
		temp[i-1] += temp[i] / 10 //进位
		temp[i] = temp[i] % 10    //去除进位的数值
	}

	// 对最高位进行处理，查看是否有进位，如果没有去除最高位的 0
	if temp[0] == 0 {
		temp = temp[1:]
	}

	// 获取结果
	res := make([]byte, len(temp))
	for i := 0; i < len(temp); i++ {
		res[i] = byte(temp[i]) + '0'
	}
	return string(res)

}
