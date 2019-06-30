package offer44

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	index int
	ans   int
}{
	{
		7,
		7,
	},
	{
		10,
		1,
	},
	{
		12,
		1,
	},
}

func Test_ac44(t *testing.T) {
	ast := assert.New(t)
	for _, ts := range tsc {
		ast.Equal(ts.ans, digitAtIndex(ts.index))
	}

}
