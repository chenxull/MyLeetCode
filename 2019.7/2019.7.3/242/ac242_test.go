package problem242

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tsc = []struct {
	s   string
	t   string
	ans bool
}{
	{
		"anagram",
		"nagaram",
		true,
	},
	{
		"rat",
		"car",
		false,
	},
}

func Test_ac242(t *testing.T) {
	ast := assert.New(t)
	for _, ts := range tsc {
		ast.Equal(ts.ans, isAnagram(ts.s, ts.t))
	}
}
