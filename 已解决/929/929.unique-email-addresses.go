package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=929 lang=golang
 *
 * [929] Unique Email Addresses
 *
 * https://leetcode.com/problems/unique-email-addresses/description/
 *
 * algorithms
 * Easy (76.96%)
 * Total Accepted:    94.6K
 * Total Submissions: 127.1K
 * Testcase Example:  '["test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"]'
 *
 * Every email consists of a local name and a domain name, separated by the @
 * sign.
 *
 * For example, in alice@leetcode.com, alice is the local name, and
 * leetcode.com is the domain name.
 *
 * Besides lowercase letters, these emails may contain '.'s or '+'s.
 *
 * If you add periods ('.') between some characters in the local name part of
 * an email address, mail sent there will be forwarded to the same address
 * without dots in the local name.  For example, "alice.z@leetcode.com" and
 * "alicez@leetcode.com" forward to the same email address.  (Note that this
 * rule does not apply for domain names.)
 *
 * If you add a plus ('+') in the local name, everything after the first plus
 * sign will be ignored. This allows certain emails to be filtered, for example
 * m.y+name@email.com will be forwarded to my@email.com.  (Again, this rule
 * does not apply for domain names.)
 *
 * It is possible to use both of these rules at the same time.
 *
 * Given a list of emails, we send one email to each address in the list.  How
 * many different addresses actually receive mails?
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input:
 * ["test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"]
 * Output: 2
 * Explanation: "testemail@leetcode.com" and "testemail@lee.tcode.com" actually
 * receive mails
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= emails[i].length <= 100
 * 1 <= emails.length <= 100
 * Each emails[i] contains exactly one '@' character.
 *
 *
 *
 */
func numUniqueEmails(emails []string) int {
	translations := make(map[string]bool, 101)
	for _, email := range emails {
		translations[clean(email)] = true
		fmt.Println(clean(email))
	}
	return len(translations)

}

//处理 ‘.’,'+','@'
func clean(email string) string {
	//根据@将 email 进行分割，后半部分不需要进行处理，主要是前半部分
	index := strings.IndexByte(email, '@')
	return deleteDot(trimPuls(email[:index])) + email[index:]
}

// 处理localName 中的‘.’
func deleteDot(localName string) string {
	return strings.Replace(localName, ".", "", -1)
}

// 处理 localName 中的 '+'
func trimPuls(localName string) string {
	index := strings.IndexByte(localName, '+')
	if index == -1 {
		return localName
	}
	fmt.Println("debug:", localName[:index])
	return localName[:index]
}

func main() {
	Input := []string{"test.emailalex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	fmt.Println(numUniqueEmails(Input))
}
