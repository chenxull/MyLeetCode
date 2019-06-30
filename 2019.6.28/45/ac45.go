package offer45

import (
	"sort"
	"strconv"
)

type intSlice []int

func (p intSlice) Less(i, j int) bool {
	return strconv.Itoa(p[i])+strconv.Itoa(p[j]) < strconv.Itoa(p[j])+strconv.Itoa(p[i])
}

func (p intSlice) Len() int {
	return len(p)
}

func (p intSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func printMinNumber(number []int) string {
	if number == nil {
		return ""
	}

	sort.Sort(intSlice(number))
	var ans string
	for _, value := range number {
		ans += strconv.Itoa(value)
	}
	//fmt.Println(ans)
	return ans
}
