package problem728

/*
 * @lc app=leetcode id=728 lang=golang
 *
 * [728] Self Dividing Numbers
 *
 * https://leetcode.com/problems/self-dividing-numbers/description/
 *
 * algorithms
 * Easy (68.85%)
 * Total Accepted:    72.4K
 * Total Submissions: 104.1K
 * Testcase Example:  '1\n22'
 *
 *
 * A self-dividing number is a number that is divisible by every digit it
 * contains.
 *
 * For example, 128 is a self-dividing number because 128 % 1 == 0, 128 % 2 ==
 * 0, and 128 % 8 == 0.
 *
 * Also, a self-dividing number is not allowed to contain the digit zero.
 *
 * Given a lower and upper number bound, output a list of every possible self
 * dividing number, including the bounds if possible.
 *
 * Example 1:
 *
 * Input:
 * left = 1, right = 22
 * Output: [1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]
 *
 *
 *
 * Note:
 * The boundaries of each input argument are 1 .
 *
 */
func selfDividingNumbers(left int, right int) []int {
	ans := make([]int, 0, right-left+1)
	//fmt.Println(ans)
	for i := left; i <= right; i++ {
		if diving(i) {
			ans = append(ans, i)
			//fmt.Println(ans)
		}
	}
	return ans
}

func diving(number int) bool {
	t := number
	for t > 0 {
		d := t % 10
		t = t / 10
		if d == 0 || number%d != 0 {
			return false
		}
	}
	return true

}
