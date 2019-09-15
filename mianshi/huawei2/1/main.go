package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	//fmt.Println(str)

	indexB :=strings.Index(str,"B")
	indexR:=strings.Index(str,"R")
	strA:=str[3:indexB-2]
	strB:=str[indexB+3:indexR-2]
	strR:=str[indexR+2:]
	//fmt.Println(strA,strB,strR)

	strNA:=strings.Split(strA,",")
	strNB:=strings.Split(strB,",")
	R,_:=strconv.Atoi(strR)

	//fmt.Println(strNA,strNB,R)
	numA:=make([]int,0,len(strNA))
	numB:=make([]int,0,len(strNB))

	for _,v:=range strNA{
		tmp,_:=strconv.Atoi(v)
		numA= append(numA,tmp)
	}

	for _,v:=range strNB{
		tmp,_:=strconv.Atoi(v)
		numB= append(numB,tmp)
	}

	//fmt.Println(numA,numB,R)

	res:=[][]int{}
	for i:=0;i<len(numA);i++{
		isAdd:=false
		for j:=0;j<len(numB);j++{
			if numA[i]<=numB[j] && numB[j] - numA[i] <=R{
				res= append(res,[]int{numA[i],numB[j]})
				isAdd =true
			}else if numA[i] <=numB[j] && !isAdd {
				res =append(res,[]int{numA[i],numB[j]})
				break
			}

		}
	}

	resStr:=""
	for _,v:=range res{
		resStr+="("+strconv.Itoa(v[0])+","+strconv.Itoa(v[1])+")"
	}
	fmt.Println(resStr)

}


