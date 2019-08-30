package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	times := 1000
	nums := make([]int, 0, times)
	for i := 0; i < times; i++ {
		val := rand.Intn(2000)
		nums = append(nums, val)
	}
	quickSort(nums, 0, len(nums)-1)

	fmt.Println(nums)
}
func quickSort(nums []int, i, j int) {
	if i >= j {
		return
	}
	p := partition(nums, i, j)
	quickSort(nums, i, p-1)
	quickSort(nums, p+1, j)

}

func partition(nums []int, l, h int) int {
	i, j, piovt := l, h, nums[l]
	for i < j {
		for piovt < nums[j] && i < j {
			j--
		}
		for nums[i] <= piovt && i < j {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i], nums[l] = nums[l], nums[i]
	return i
}
