package problem147

/*
 * @lc app=leetcode id=147 lang=golang
 *
 * [147] Insertion Sort List
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

// 需要理解插入排序的过程
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummyhead := &ListNode{Next: head}

	// cur 指向当前节点，p是待插入的节点 指向 cur.next
	cur := head

	for cur != nil && cur.Next != nil {
		// 待插入节点
		p := cur.Next
		if cur.Val < p.Val {
			cur = cur.Next
			continue
		}

		// 保留后续节点信息，在进行遍历寻找
		cur.Next = p.Next
		// 重头开始遍历整个有序的链表，寻找插入位置
		// 因为需要待插入位置的前置节点，所以需要保存二个节点信息
		pre, next := dummyhead, dummyhead.Next

		for next.Val < p.Val {
			// 没找到合适位置
			pre = pre.Next
			next = next.Next
		}
		// 找到合适位置
		pre.Next = p
		p.Next = next
	}
	return dummyhead.Next
}
