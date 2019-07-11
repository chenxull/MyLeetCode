package problem20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tsc = []struct {
	input string
	ans   bool
}{
	{
		"()",
		true,
	},
	{
		"()[]{}",
		true,
	},
	{
		"([)]",
		false,
	},
}

func Test_ac20(t *testing.T) {
	ast := assert.New(t)
	for _, ts := range tsc {
		ast.Equal(ts.ans, isValid(ts.input))
	}
}
