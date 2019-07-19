package offer43

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	n   int
	ans int
}{
	{
		12,
		5,
	}, {
		534,
		214,
	},
}

func Test_ac43(t *testing.T) {
	ast := assert.New(t)
	for _, ts := range tsc {
		ast.Equal(ts.ans, numberOf1Between1AndN(ts.n))
	}

}
