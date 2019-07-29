package main

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/rand"
)

func main() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.") // 1
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done: // 用来控制 goroutine 的终止
					return
				}
			}
		}()

		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)
	time.Sleep(1 * time.Second)
}
