package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
555503,772867,756893,339138,399447,
40662,859037,628085,855723,974471,
599631,636153,581541,174761,948135,
411485,554033,858627,402833,546467,
574367,360461,566480,755523,222921,
164287,420256,40043,977167,543295,944841,917125,331763,188173,353275,175757,950417,284578,617621,546561,
913416,258741,260569,630583,252845;10

*/
func main() {
	var str string
	fmt.Scan(&str)
	// fmt.Println(str)
	nums := strings.Split(str, ",")
	temp := nums[len(nums)-1]
	lastStr := strings.Split(temp, ";")
	N := lastStr[1]

	nums = nums[:len(nums)-1]
	nums = append(nums, lastStr[0])

	// 用二个数组分别来保存偶数和奇数
	numsOdd := make([]int, 0, len(nums))
	numsEven := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		c := nums[i]
		num, _ := strconv.Atoi(c)
		if num%2 == 0 {
			numsEven = append(numsEven, num)
		} else {
			numsOdd = append(numsOdd, num)
		}
	}

	sort.Slice(numsEven, func(i, j int) bool {
		return numsEven[i] > numsEven[j]
	})

	sort.Slice(numsOdd, func(i, j int) bool {
		return numsOdd[i] > numsOdd[j]
	})
	// fmt.Println(numsEven)
	// fmt.Println(numsOdd, N)
	NN, _ := strconv.Atoi(N)
	res := ""
	if len(numsEven) > NN {
		// fmt.Println(numsEven[:NN])

		for i := 0; i < NN; i++ {
			if i == 0 {
				res += strconv.Itoa(numsEven[0])
			} else {
				res += "," + strconv.Itoa(numsEven[i])

			}
		}

		fmt.Println(res)
		return
	}
	rest := NN - len(numsEven)

	for i := 0; i < len(numsEven); i++ {
		if i == 0 {
			res += strconv.Itoa(numsEven[i])
			continue
		}
		res += "," + strconv.Itoa(numsEven[i])
	}

	for i := 0; i < rest; i++ {
		res += "," + strconv.Itoa(numsOdd[i])
	}
	fmt.Println(res)
	return
}
