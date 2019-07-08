package problem125

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tsc = []struct {
	input string
	ans   bool
}{
	{
		"A man, a plan, a canal: Panama",
		true,
	},
	{
		"race a car",
		false,
	},
}

func Test_ac125(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tsc {
		ast.Equal(tc.ans, isPalindrome(tc.input))
	}
}
