package problem763

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tsc = []struct {
	S   string
	ans []int
}{
	{
		"ababcbacadefegdehijhklij",
		[]int{9, 7, 8},
	},
}

func Test_ac763(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tsc {
		ast.Equal(tc.ans, partitionLabels(tc.S))
	}
}
