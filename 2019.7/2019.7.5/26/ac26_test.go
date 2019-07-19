package problem26

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tsc = []struct {
	nums []int
	ans  int
}{
	{
		[]int{1, 1, 2},
		2,
	}, {
		[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		5,
	},
}

func Test_ac26(t *testing.T) {
	ast := assert.New(t)
	for _, ts := range tsc {
		ast.Equal(ts.ans, removeDuplicates(ts.nums))
	}

}
