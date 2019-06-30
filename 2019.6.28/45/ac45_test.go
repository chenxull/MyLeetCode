package offer45

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	input []int
	ans   string
}{
	{
		[]int{3, 32, 321},
		"321323",
	},
}

func Test_ac45(t *testing.T) {
	ast := assert.New(t)
	for _, ts := range tsc {
		ast.Equal(ts.ans, printMinNumber(ts.input))
	}

}
