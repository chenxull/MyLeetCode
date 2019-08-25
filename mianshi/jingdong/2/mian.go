package main

import (
	"fmt"
	"unicode"
)

func main() {
	var n int
	fmt.Scan(&n)
	var str string
	fmt.Scan(&str)

	size := len(str)
	Rune := []rune(str)
	nums := make([]int, 0, size)
	for i := 0; i < size; i++ {
		if unicode.IsLower(Rune[i]) {
			nums = append(nums, 0)
		}

		if unicode.IsUpper(Rune[i]) {
			nums = append(nums, 1)
		}
	}
	count := 0
	isCap := 0
	for i := range nums {
		if isCap == 0 {
			if i == size-1 {
				if nums[i] == 0 {
					count++
					break
				}
				if nums[i] == 1 {
					count += 2
					break
				}
			}
			if nums[i] == 0 {
				count++
				continue
			}
			if nums[i] == 1 && nums[i+1] == 0 {
				count += 2
				continue
			}

			if nums[i] == 1 && nums[i+1] == 1 {
				count += 2
				isCap = 1
				continue
			}
		}
		if isCap == 1 {
			if i == size-1 {
				if nums[i] == 0 {
					count += 2
					break
				}
				if nums[i] == 1 {
					count++
					break
				}
			}
			if nums[i] == 1 {
				count++
				continue
			}
			if nums[i] == 0 && nums[i+1] == 1 {
				count += 2
				continue
			}
			if nums[i] == 0 && nums[i+1] == 0 {
				count += 2
				continue
			}
		}
	}
	fmt.Println(count)
}
