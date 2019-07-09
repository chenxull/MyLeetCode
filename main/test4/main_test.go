package main

import "testing"

var tsc = []struct {
	b     []byte
	times int
	ans   string
}{
	{
		[]byte("a"),
		5,
		"aaaaa",
	},
}

func Benchmark_main(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tsc {
			makeString1(tc.b[0], tc.times)
		}
	}
}

func Benchmark_main1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tsc {
			makeString2(tc.b[0], tc.times)
		}
	}
}
