package offer38

import "fmt"

func permutation(arrayChar []byte, start int) {
	if len(arrayChar) <= 1 {
		return
	}

	if start == len(arrayChar)-1 {
		for i := 0; i < len(arrayChar); i++ {
			fmt.Printf("%c", arrayChar[i])
		}
		fmt.Println()
	} else {
		for i := start; i < len(arrayChar); i++ {
			swap(arrayChar, i, start)
			permutation(arrayChar, start+1)
			swap(arrayChar, i, start)
		}
	}

}

func swap(array []byte, a, b int) {
	temp := array[a]
	array[a] = array[b]
	array[b] = temp
}
