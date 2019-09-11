package main

import "fmt"

func main(){
	var N int
	fmt.Scan(&N)

	res:=0
	for i:=1;i<=N;i++{
		target :=N
		cur :=i
		for target > 0{
			target-=cur
			cur++
		}
		if target == 0{
			res++
		}
	}
	fmt.Println(res)
}