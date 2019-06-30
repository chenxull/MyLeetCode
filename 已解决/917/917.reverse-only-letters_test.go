package problem917

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  string
	output string
}{
	{
		"ab-cd",
		"dc-ba",
	},
	{
		"j-Ih-gfE-dCba",
		"a-bC-dEf-ghIj",
	},
	{
		"Test1ng-Leet=code-Q!",
		"Qedo1ct-eeLg=ntse-T!",
	},
}

func Test_reverseOnlyLetters(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Println("输入:", tc.input, tc.output)
		ast.Equal(tc.output, reverseOnlyLetters(tc.input))
		fmt.Println("输出", reverseOnlyLetters(tc.input))
	}
}
func Benchmark_reverseOnlyLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reverseOnlyLetters(tc.input)
		}
	}
}
