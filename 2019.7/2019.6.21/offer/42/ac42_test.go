package offer42

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	data []int
	ans  int
}{
	{
		[]int{1, -2, 3, 10, -4, 7, 2, -5},
		18,
	},
}

func Test_ac42(t *testing.T) {
	ast := assert.New(t)
	for _, ts := range tsc {
		ast.Equal(ts.ans, FindGreatestSumOfSubArray(ts.data))
	}
}
