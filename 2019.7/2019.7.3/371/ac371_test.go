package problem371

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	a int
	b int
}

var tsc = []struct {
	input
	ans int
}{
	{
		input{
			1,
			2,
		},
		3,
	},
}

func Test_ac371(t *testing.T) {
	ast := assert.New(t)
	for _, ts := range tsc {
		ast.Equal(ts.ans, getSum(ts.a, ts.b))
	}
}
