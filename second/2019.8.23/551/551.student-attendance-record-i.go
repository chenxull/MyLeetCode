package problem551

// import "fmt"

/*
 * @lc app=leetcode id=551 lang=golang
 *
 * [551] Student Attendance Record I
 */
func checkRecord(s string) bool {
	if len(s) == 0 {
		return false
	}
	countA := 0
	countL := 0
	for _, c := range s {
		if c == 'A' {
			countA++
		}
		if c == 'L' {
			countL++
		}
	}

	if countA > 1 || countL > 2 {
		return false
	}
	return true
}

// func main() {
// 	s := "PPALLL"
// 	fmt.Println(checkRecord(s))
// }
