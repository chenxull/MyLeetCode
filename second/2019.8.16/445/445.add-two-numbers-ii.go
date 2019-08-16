package problem445

/*
 * @lc app=leetcode id=445 lang=golang
 *
 * [445] Add Two Numbers II
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 先将链表转化为切片的形式
	s1 := make([]int, 0, 64)
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}

	s2 := make([]int, 0, 64)
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	head := &ListNode{}
	sum, carry := 0, 0
	// 从切片的最后一位开始处理，head 表示当前位置的值，front 表示进位的值
	for len(s1) != 0 || len(s2) != 0 {
		sum = carry
		if len(s1) != 0 {
			sum += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}

		if len(s2) != 0 {
			sum += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}

		carry = sum / 10
		// 这个链表构成的思想，一般都应用在从尾部开始，使用二个指针，一个指向需要处理的元素这里是 head
		// 另一个指针为其前一个位置的指针，这里是 front
		head.Val = sum % 10
		// front 保存进位信息
		front := &ListNode{Val: sum / 10}
		front.Next = head
		head = front
	}

	// 因为 head 最后指向的是 front，可能存在 front 为0 的情况，所以需要判断一下 head 的值
	if head.Val == 0 {
		return head.Next
	}
	return head
}
