package offer39

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	numbers []int
	ans     int
}{
	{
		[]int{0, 1, 2, 3, 4, 5, 6, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		1,
	}, {
		[]int{2, 34, 54, 3, 1, 5, 1, 3, 5, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		2,
	},
}

func Test_ac39(t *testing.T) {
	ast := assert.New(t)
	for _, ts := range tsc {
		ast.Equal(ts.ans, moreThanHalfNum(ts.numbers), "相等")
	}
}
