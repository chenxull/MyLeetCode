package problem401

import "fmt"

/*
 * @lc app=leetcode id=401 lang=golang
 *
 * [401] Binary Watch
 */
func readBinaryWatch(num int) []string {
	res := []string{}
	//  用来标识哪些灯是亮的
	leds := make([]bool, 10)
	dfs(num, 0, leds, &res)
	return res
}
func dfs(num int, start int, leds []bool, res *[]string) {
	//  回溯终止条件
	if num == 0 {
		// 分钟的灯保存咋前 6 位， 小时的灯保存在后 4 为。通过切片的方式区分
		m, h := get(leds[:6]), get(leds[6:])
		if h < 12 && m < 60 {
			*res = append(*res, fmt.Sprintf("%d:%02d", h, m))
			fmt.Println(*res)
		}
		return
	}

	//  len(leds) - n +1 表示 有多少位
	for i := start; i < len(leds)-num+1; i++ {
		leds[i] = true
		fmt.Println(leds)
		dfs(num-1, i+1, leds, res)
		leds[i] = false
	}
}

var bs = []int{1, 2, 4, 8, 16, 32}

//  获取时间
func get(leds []bool) int {
	sum := 0
	//  遍历亮灯记录，如果为 ture 则表示这个led 是亮的，根据其位置信息，判断出其数值是多少
	for i := 0; i < len(leds); i++ {
		if leds[i] {
			sum += bs[i]
		}
	}
	return sum
}

// func main() {

// 	fmt.Println(readBinaryWatch(2))
// }
