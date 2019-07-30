package problem731

import "sort"

/*
 * @lc app=leetcode id=731 lang=golang
 *
 * [731] My Calendar II
 */

// MyCalendarTwo is
type MyCalendarTwo struct {
	events [][]int
}

// Constructor is
func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

// Book is
func (m *MyCalendarTwo) Book(start int, end int) bool {
	// s用一个三元组表示订阅信息
	e := []int{start, end, 1}

	n := len(m.events)
	// 当日历为空时
	if n == 0 {
		m.events = append(m.events, e)
		return true
	}

	//  i,j 分别表示 新的订阅信息 可以插入的位置
	i := sort.Search(n, func(i int) bool {
		return e[0] < m.events[i][1]
	})

	j := sort.Search(n, func(i int) bool {
		return e[1] < m.events[i][0]
	})
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
