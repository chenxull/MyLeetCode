package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var m int
	fmt.Scan(&n)
	length := pow(n)
	numsInput := make([]int, 0, length)
	// 获取输入
	for i := 0; i < int(length); i++ {
		temp := 0
		fmt.Scan(&temp)
		numsInput = append(numsInput, temp)
	}

	if len(numsInput) == 0 {
		return
	}

	// 获取 m
	fmt.Scan(&m)
	// nums 中存储的是分组的大小
	nums := make([]int, 0, m)
	for i := 0; i < m; i++ {
		temp := 0
		fmt.Scan(&temp)
		nums = append(nums, temp)
	}

	if len(nums) == 0 {
		return
	}

	/*
		2
		2 1 4 3
		4
		1 2 0 2
	*/
	// 开始处理
	for i := 0; i < m; i++ {
		count := 0
		var q int
		q = nums[i]
		size := pow(q)

		if size == 0 {
			count = calculate(numsInput)
			fmt.Println(count)
		} else if size == length {
			reverse(numsInput, 0, size)
			count = calculate(numsInput)
			fmt.Println(count)
		} else {
			// 对于需要分组的情况处理
			for j := 0; j <= n; j += size {
				reverse(numsInput, j, j+size)
			}
			count = calculate(numsInput)
			fmt.Println(count)
		}

	}
}

// 反转
func reverse(nums []int, i, j int) {
	j = j - 1
	for i <= j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// 计算逆序对的数量
func calculate(nums []int) int {
	count := 0
	size := len(nums)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if nums[i] > nums[j] {
				count++
			}
		}
	}
	return count
}
func pow(a int) int {
	var ans int
	ans = int(math.Pow(2, float64(a)))
	return ans
}
