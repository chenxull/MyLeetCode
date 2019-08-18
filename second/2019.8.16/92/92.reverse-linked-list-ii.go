package problem92

/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
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

// 1.找到需要切割的位置，断开
// 2.反转链表，返回需要的节点信息
// 3.拼接链表
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}

	dummyHead := &ListNode{Next: head}
	m++
	n++
	pre, node, next := split(dummyHead, m, n)

	first, end := reverse(node)
	pre.Next = first
	end.Next = next

	return dummyHead.Next
}

// 需要从 dummyhead 开始，加入 m = 1，需要使用 dummyhead 保存首节点的信息
func split(head *ListNode, m, n int) (Pre, node, next *ListNode) {
	i := 1
	for head != nil {
		if i == m-1 {
			Pre = head
			node = head.Next
		}

		// next 保存的是n+1 位置的节点信息，用于后续的连接
		if i == n {
			next = head.Next
			head.Next = nil

		}
		head = head.Next
		i++
	}
	return
}

// 使用递归的方式
func reverse(head *ListNode) (first, end *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}
	var e *ListNode

	first, e = reverse(head.Next)
	e.Next = head
	end = head
	return
}
