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

//  从原链表中截取需要的部分，对截取返回进行反转，然后在接入到原来的链表中
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}

	//  头结点
	headPre := &ListNode{}
	headPre.Next = head
	m++
	n++

	//  将需要反转的部分切割出来
	mPre, mNode, mNext := split(headPre, m, n)

	// 对切割出来的部分反转
	h, e := reverse(mNode)

	mPre.Next = h
	e.Next = mNext

	return headPre.Next
}

func split(head *ListNode, m, n int) (mPre, mNode, mNext *ListNode) {
	i := 1
	for head != nil {
		//  需要保留前序节点
		if i == m-1 {
			// 保存之前的链表指针
			mPre = head
			mNode = head.Next
		}
		if i == n {
			//  保留后续的链表指针
			mNext = head.Next
			//  截断
			head.Next = nil
		}
		head = head.Next
		i++
	}
	return
}

//  递归逆序链表，需要继续研究
func reverse(head *ListNode) (h, e *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}

	var end *ListNode
	h, end = reverse(head.Next)
	end.Next = head
	e = head
	return
}
