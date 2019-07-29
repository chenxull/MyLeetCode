package main

import (
	"fmt"
	"time"
)

func main() {

	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	strs := make(chan string, 100)
	strs <- "chenxu"
	terminated := doWork(done, nil)

	// 在这里我们创建另一个goroutine，2秒后就会取消doWork中产生的goroutine。
	go func() {
		// Cancel the operation after 2 second
		time.Sleep(2 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()
	// 这是我们在main goroutine中调用doWork函数返回结果的地方。
	<-terminated
	fmt.Println("Done.")

}
