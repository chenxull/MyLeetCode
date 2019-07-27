package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func quickSort(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}

	p := partition1(nums, lo, hi)
	quickSort(nums, lo, p-1)
	quickSort(nums, p+1, hi)
}

func quickSortGo(nums []int, lo, hi int, done chan struct{}, depth int) {
	if lo >= hi {
		done <- struct{}{}
		return
	}

	// 控制并发数量
	depth--
	p := partition1(nums, lo, hi)
	if depth > 0 {
		childDone := make(chan struct{}, 2)
		go quickSortGo(nums, lo, p-1, childDone, depth)
		go quickSortGo(nums, p+1, hi, childDone, depth)
		// 阻塞二个
		<-childDone
		<-childDone
	} else {
		quickSort(nums, lo, p-1)
		quickSort(nums, p+1, hi)
	}
	done <- struct{}{}

}

func partition1(nums []int, l, h int) int {

	i, j, pivot := l, h, nums[l]
	for i < j {
		//  必须从大的数字开始计算
		for pivot < nums[j] && i < j {
			j--
		}
		for nums[i] <= pivot && i < j {
			i++
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[l] = nums[l], nums[i]
	return i
}

func partition2(nums []int, lo, hi int) int {
	pivot := nums[hi]
	//[0,i]区间存储的都是小于 pivot 的元素
	i := lo - 1
	// 从低位开始遍历,遍历完之后，[0,i] 都是小于 pivot 元素
	for j := lo; j < hi; j++ {
		// 如果当前位的值小于 pivot，将其与 i+1 位置的元素交换位置。这样就确保[0,i]都是小于 pivot 元素
		if nums[j] < pivot {
			i++
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	nums[i+1], nums[hi] = nums[hi], nums[i+1]
	return i + 1

}

func main() {
	rand.Seed(time.Now().UnixNano())
	testData1, testData2, testData3, testData4 := make([]int, 0, 100000000), make([]int, 0, 100000000), make([]int, 0, 100000000), make([]int, 0, 100000000)
	times := 1000000
	for i := 0; i < times; i++ {
		val := rand.Intn(2000000)
		testData1 = append(testData1, val)
		testData2 = append(testData2, val)
		testData3 = append(testData3, val)
		testData4 = append(testData4, val)
	}

	start := time.Now()
	quickSort(testData1, 0, len(testData1)-1)
	fmt.Println("signle goroutine:", time.Now().Sub(start))
	if !sort.IntsAreSorted(testData1) {
		fmt.Println("wrong quick_sort implementation")
	}

	done := make(chan struct{})
	start = time.Now()
	go quickSortGo(testData2, 0, len(testData2)-1, done, 15)
	<-done
	fmt.Println("mutil:", time.Now().Sub(start))
	if !sort.IntsAreSorted(testData2) {
		fmt.Println("wrong quicksortgo implementation ")
	}

}
