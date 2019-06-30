package Problem824

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	S   string
	ans string
}{
	{
		"I speak Goat Latin",
		"Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
	},
	{
		"The quick brown fox jumped over the lazy dog",
		"heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
	},
}

func Test_toGoatLatin(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Println("输入", tc.S)
		ast.Equal(tc.ans, toGoatLatin(tc.S), "输入:", tc.S)
		fmt.Println("输出", toGoatLatin(tc.S))
	}
}

func Benchmark_toGoatLatin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			toGoatLatin(tc.S)
		}
	}
}
