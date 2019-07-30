package problem729

/*
 * @lc app=leetcode id=729 lang=golang
 *
 * [729] My Calendar I
 */
type event struct {
	start, end int
	next       *event
}

// MyCalendar is
type MyCalendar struct {
	head *event
}

// Constructor is
func Constructor() MyCalendar {
	var head *event
	return MyCalendar{head: head}
}

// Book is  使用链表保存订阅信息 。插入日程信息时，插入的位置有 4 中可能的状态
func (m *MyCalendar) Book(start int, end int) bool {
	// 新的订阅信息
	e := &event{start: start, end: end}
	// 状态 1 日历为空
	// 状态 2 新的订阅信息在日历中排在第一位
	if m.head == nil ||
		e.end <= m.head.start {
		//  移动链表
		e.next, m.head = m.head, e
		return true
	}

	// 寻找订阅信息插入的位置,设置一个 dummy 节点
	dummy := &event{next: m.head}
	// 跳过所有在 e 开始之前的事件
	for dummy.next != nil && dummy.next.end <= e.start {
		dummy = dummy.next
	}

	// 状态 3 订阅信息放在日历最后的位置
	// 状态 4 订阅信息在 夹在二个订阅信息之间
	if dummy.next == nil ||
		e.end <= dummy.next.start {
		e.next, dummy.next = dummy.next, e
		return true
	}
	return false

}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
