package main

import "fmt"

/*
	 set:将指定位置设置为 1
     clear:将指定位置设置为 0
	 get：查找指定位置的值
     实现自动扩容
*/

func main() {
	a := 64
	a = a >> 6
	fmt.Println(a)
}
