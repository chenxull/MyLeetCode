package problem35

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tsc = []struct {
	nums   []int
	target int
	ans    int
}{
	{
		[]int{1, 3, 5, 6},
		5,
		2,
	},
	{
		[]int{1, 3, 5, 6},
		2,
		1,
	},
	{
		[]int{1, 3, 5, 6},
		7,
		4,
	},
	{
		[]int{1, 3, 5, 6},
		0,
		0,
	},
}

func Test_ac35(t *testing.T) {
	ast := assert.New(t)

	for _, ts := range tsc {
		ast.Equal(ts.ans, searchInsert(ts.nums, ts.target))
	}
}
