package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var str string
	fmt.Scan(&str)
	is_capslocks := false

	count := 0
	Rune := []rune(str)
	for i := 0; i < n; i++ {
		if Rune[i] >= 'A' && Rune[i] <= 'Z' {
			if is_capslocks == true {
				count++
			} else {
				if Rune[i+1] >= 'A' && Rune[i+1] <= 'Z' {
					is_capslocks = !is_capslocks
				}
				count += 2
			}
		} else {
			if is_capslocks == false {
				count++
			} else {
				if Rune[i+1] >= 'a' && Rune[i+1] <= 'z' {
					is_capslocks = !is_capslocks
				}
				count += 2
			}
		}
	}
	fmt.Println(count)

}
