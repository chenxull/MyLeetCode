package main

import (
	"fmt"
	"time"
)

func main() {
	// 你可以将任意数量的通道组合到单个通道中，
	// 只要任何作为组件的通道关闭或被写入，整个通道就会关闭。
	var or func(channels ...<-chan interface{}) <-chan interface{}

	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}

		orDone := make(chan interface{})
		go func() {
			defer close(orDone)
			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):

				}
			}
		}()
		return orDone
	}

	sig := func(after time.Duration) <-chan interface{} { //1
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now() //2
	<-or(sig(2*time.Hour), sig(5*time.Minute), sig(1*time.Second), sig(1*time.Hour), sig(1*time.Minute))
	fmt.Printf("done after %v", time.Since(start)) //3
}
