package problem804

/*
 * @lc app=leetcode id=804 lang=golang
 *
 * [804] Unique Morse Code Words
 */

var table = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	if len(words) == 0 {
		return 0
	}

	// 使用 hashmap 保存结果，这样就可以避免重复计算
	res := make(map[string]bool, len(words))
	for _, w := range words {
		str := ""
		for i := 0; i < len(w); i++ {
			str += table[w[i]-'a']
		}
		res[str] = true
	}

	return len(res)
}

// func main() {
// 	str := []string{"gin", "zen", "gig", "msg"}
// 	fmt.Println(uniqueMorseRepresentations(str))
// }
