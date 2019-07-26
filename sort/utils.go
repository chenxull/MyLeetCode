package utils

import (
	"math/rand"
	"time"
)

// ArrayRand 生成随机数
func ArrayRand(n int) []int {
	rand.Seed(time.Now().UnixNano())

	array := make([]int, n)
	for i := range array {
		array[i] = rand.Intn(n)
	}
	return array
}

//Swap 交换
func Swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
