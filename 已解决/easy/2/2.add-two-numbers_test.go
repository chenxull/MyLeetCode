package problem2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type para struct {
	one *ListNode
	two *ListNode
}

type ans struct {
	as *ListNode
}

type question struct {
	p para
	a ans
}

func makeListNode(numbers []int) *ListNode {
	if numbers == nil {
		return nil
	}
	res := &ListNode{
		Val: numbers[0],
	}
	tmp := res
	for i := 1; i < len(numbers); i++ {
		Next = &ListNode{
			Val: numbers[i],
		}
		tmp = Next
	}
	return res
}

func Test_addTwoNumbers(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		{
			p: para{
				one: makeListNode([]int{1, 2, 3}),
				two: makeListNode([]int{3, 4, 6}),
			},
			a: ans{
				as: makeListNode([]int{4, 6, 9}),
			},
		},
		{
			p: para{
				one: makeListNode([]int{4, 5, 1}),
				two: makeListNode([]int{6, 2, 6}),
			},
			a: ans{
				as: makeListNode([]int{0, 8, 7}),
			},
		},
		{
			p: para{
				one: makeListNode([]int{5, 6, 1}),
				two: makeListNode([]int{6, 7, 2}),
			},
			a: ans{
				as: makeListNode([]int{1, 4, 4}),
			},
		},
		{
			p: para{
				one: makeListNode([]int{1, 3, 3}),
				two: makeListNode([]int{3, 4, 6}),
			},
			a: ans{
				as: makeListNode([]int{4, 7, 9}),
			},
		},
	}

	for _, q := range qs {
		p, a := q.p, q.a
		ast.Equal(a.as, addTwoNumbers(p.one, p.two), "输入:%v", p)
	}
}
