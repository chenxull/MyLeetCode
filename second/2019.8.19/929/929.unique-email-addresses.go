package main

import (
	"fmt"
	"strings"
)

// package problem929

// import (
// 	"fmt"
// 	"strings"
// )

/*
 * @lc app=leetcode id=929 lang=golang
 *
 * [929] Unique Email Addresses
 */

//  主要是对local name 的操作
func numUniqueEmails(emails []string) int {
	if len(emails) == 0 {
		return 0
	}

	res := make(map[string]bool, len(emails))
	for _, email := range emails {
		// 切分获取 localName
		name := strings.Split(email, "@")
		localName := name[0]
		// 对 localName 进行处理
		localName = clean(localName)
		realName := localName + "@" + name[1]
		fmt.Println(realName)
		res[realName] = true
	}
	return len(res)
}

func clean(localName string) string {
	return replacedot(localName)
}

// 取代 .
func replacedot(localName string) string {
	localName = strings.Replace(localName, ".", "", -1)
	localName = deletePlus(localName)

	return localName
}
func deletePlus(localName string) string {
	index := strings.IndexByte(localName, '+')
	if index == -1 {
		return localName
	}
	return localName[:index]
}

func main() {
	emails := []string{"test.email+alex@leetcode.com", "test.email.leet+alex@code.com"}
	fmt.Println(numUniqueEmails(emails))
}
