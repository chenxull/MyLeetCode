package problem438

/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 */
func findAnagrams(s string, p string) []int {
	//  将 p 的值放入 hash 表中，每次移动窗口时，判断 hash表即可

	left, right, n, count := 0, 0, len(s), 0
	pLen := len(p)
	hash := [256]int{}
	// res := make([]int, 0, len(s))
	res := []int{}
	//  将 p 的值，建立 hash 表
	for i := 0; i < pLen; i++ {
		hash[p[i]]++
	}

	for right < n {
		//  存在 p 中的元素
		if hash[s[right]] > 0 {
			hash[s[right]]--
			count++
			right++
		} else {
			//  不存在，为了维护 hash 表的完成性，需要将刚刚最左边的加上，并向右移动
			hash[s[left]]++
			count--
			left++
		}

		if count == pLen {
			res = append(res, left)
		}
	} 
	return res

}
