package offer40

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	data []int
	k    int
	ans  []int
}{
	{
		[]int{4, 3, 5, 1, 2, 5, 7, 9},
		3,
		[]int{1, 2, 3},
	},
}

func Test_ac40(t *testing.T) {
	ast := assert.New(t)
	for _, ts := range tsc {
		ast.Equal(ts.ans, getLeastNumbers(ts.data, ts.k))
	}
}
