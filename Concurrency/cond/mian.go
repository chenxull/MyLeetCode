package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		fmt.Println("Remove from queue: ", queue[0])
		queue = queue[1:]
		c.L.Unlock()
		//  这个信号 会释放 main goroutine 的阻塞
		c.Signal() 
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		// 在这里我们检查队列的长度，以确认什么时候需要等待。由于removeFromQueue是异步的，
		// for不满足时才会跳出，而if做不到重复判断，这一点很重要。
		for len(queue) == 3 {
			// 调用Wait，这将阻塞main goroutine，直到接受到信号。
			c.Wait()
		}
		fmt.Println("Adding to queue :", i)
		queue = append(queue, i)
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}
