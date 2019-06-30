package problem942

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input string
}{

	{
		"IDID",
	},

	{
		"III",
	},

	{
		"DDI",
	},
}

func isCorrect(s string, a []int) bool {
	for i, value := range s {
		switch value {
		case 'I':
			if !(a[i] < a[i+1]) {
				return false
			}
		case 'D':
			if !(a[i] > a[i+1]) {
				return false
			}
		default:
			panic("ERROR input ")
		}
	}
	return true
}

func Test_diStringMatch(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		a := diStringMatch(tc.input)
		ast.True(isCorrect(tc.input, a))
	}
}
