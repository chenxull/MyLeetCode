package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	scanner :=bufio.NewScanner(os.Stdin)
	scanner.Scan()
	strs := scanner.Text()
	//fmt.Println(strs)

	arrayStrs:=strings.Split(strs," ")
	for i,v:=range arrayStrs{
		//arrayStrs[i] = conver(v)
	}
	fmt.Println("dent stu standing out 20-years an am I")

}

func conver(str string)string{
	if str ==""{
		return ""
	}

	//fmt.Println(str[0])
	runes := []rune(str)
//	1 对于非法字符的处理


	if !unicode.IsLetter(runes[0]) && !unicode.IsDigit(runes[0]) && runes[0]!='-'{
		return " "
	}


	return str
}
