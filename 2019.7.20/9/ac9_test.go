package problem9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tsc = []struct {
	input int
	ans   bool
}{
	{
		121,
		true,
	},
	{
		-121,
		false,
	},
	{
		10,
		false,
	},
}

func Test_ac9(t *testing.T) {
	ast := assert.New(t)

	for _, ts := range tsc {
		ast.Equal(ts.ans, isPalindrome(ts.input))
	}
}
