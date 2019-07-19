package problem171

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tsc = []struct {
	input string
	ans   int
}{
	{
		"A",
		1,
	},
	{
		"AB",
		28,
	},
	{
		"ZY",
		701,
	},
}

func Test_ac171(t *testing.T) {
	ast := assert.New(t)
	for _, ts := range tsc {
		ast.Equal(ts.ans, titleToNumber(ts.input))
	}
}
