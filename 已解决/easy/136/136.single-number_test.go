package problem136

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	input  []int
	output int
}{
	{
		input:  []int{2, 2, 1},
		output: 1,
	},
	{
		input:  []int{4, 1, 2, 1, 2},
		output: 4,
	},
}

func Test_ok(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tsc {
		ast.Equal(tc.output, singleNumber(tc.input))
	}
}
