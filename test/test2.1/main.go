package main

import (
	"fmt"
	"unsafe"
)

type Programmer struct {
	name     string
	language string
}

func main() {

	s := make([]int, 9, 9)

	// &s => pointer => uintptr => pointer => *int => int
	var len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	var cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println("len:", len, "cap:", cap)

	p := Programmer{"chenxu", "go"}
	fmt.Println(p)

	name := (*string)(unsafe.Pointer(&p))
	*name = "chenxull"

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
	*lang = "c++"
	fmt.Println(p)

}
