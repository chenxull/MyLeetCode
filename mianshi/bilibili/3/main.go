package main

import (
	"fmt"
	"strings"
)

func main() {
	var kvpd,kvd,str string
	fmt.Scan(&kvpd,&kvd,&str)

	strs := strings.Split(str,kvpd)
	size :=len(strs)

	fmt.Println(size)
	for _,v:=range  strs{
		kv:=strings.Split(v,kvd)
		fmt.Println(kv[0],kv[1])
	}
}
