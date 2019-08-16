package problem328

/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
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

// 直接构建二个链表，根据奇偶位将数据放入到对应的链表中。最后合并链表
func oddEvenList(head *ListNode) *ListNode {
	// 头结点，为空
	oddhead := &ListNode{}
	evenhead := &ListNode{}

	odd := oddhead
	even := evenhead
	count := 1

	for head != nil {
		if count%2 != 0 {
			odd.Next = head
			odd = odd.Next
		} else {
			even.Next = head
			even = even.Next
		}
	}

	even.Next = nil
	odd.Next = evenhead.Next

	head = oddhead.Next
	return head
}
