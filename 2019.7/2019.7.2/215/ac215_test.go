package problem215

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	nums []int
	k    int
	ans  int
}{

	{
		[]int{3, 3, 3, 3, 3, 3, 3, 3, 3},
		1,
		3,
	},

	{
		[]int{3, 2, 1, 5, 6, 4},
		2,
		5,
	},
	{
		[]int{3, 2, 1, 5, 6, 4},
		2,
		5,
	},

	// 可以有多个 testcase
}

func Test_ac215(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, findKthLargest(tc.nums, tc.k))
	}
}
