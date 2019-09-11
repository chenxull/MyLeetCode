package main

import "fmt"

// 数字生成器
func xrange() chan int {
	ch := make(chan int)

	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter(in chan int, number int) chan int {
	out := make(chan int)
	go func() {
		for {
			i := <-in

			if i%number != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	const max = 10000
	nums := xrange()
	number := <-nums

	for number < max { //终止条件
		fmt.Println(number)
		// 过滤 number 的倍数
		nums := filter(nums, number)
		number = <-nums
	}
}
