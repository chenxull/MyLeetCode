package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	num := amout(n, m)
	ans := (num * amout(m*2-1, m-1)) % 1000000007
	fmt.Println(ans)

}

func amout(a, b int) int {
	da := 1
	x := 1
	for i := 0; i < b; i++ {
		da = da * a
		a--
	}
	for ; b > 0; b-- {
		x = x * b
	}
	return da / x
}
