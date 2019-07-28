package main

import "fmt"

func makeString1(b byte, n int) string {
	res := make([]byte, 0, n)
	for i := range res {
		res[i] = b
	}
	return string(res)
}

func makeString2(b byte, n int) string {
	res := ""
	for i := 0; i < n; i++ {
		res += string(b)
	}
	return res
}

func main() {
	b := []byte("s")
	res1 := makeString1(b[0], 3)
	res2 := makeString2(b[0], 3)
	fmt.Println("res1:", res1, "  res2:", res2)
}
