package problem19

/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用双指针，其中一个先移动 n 个位置，另外一个在开始 遍历
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	dummyhead := &ListNode{Next: head}

	slow, fast := dummyhead, dummyhead

	for fast.Next != nil && n > 0 {
		n--
		fast = fast.Next
	}

	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}

	delNode := slow.Next
	slow.Next = delNode.Next

	return dummyhead.Next

}
