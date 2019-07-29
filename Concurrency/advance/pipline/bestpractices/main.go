package main

import (
	"fmt"
	"time"
)

func main() {

	generator := func(done <-chan interface{}, integets ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integets {
				select {
				case <-done:
					fmt.Println("close generator")
					return
				case intStream <- i:

				}
			}
		}()
		return intStream
	}

	multiply := func(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
					fmt.Println("close multiply")

					return
				case multipliedStream <- i * multiplier:

				}
			}
		}()
		return multipliedStream
	}

	add := func(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for i := range intStream {
				select {
				case <-done:
					fmt.Println("close add")
					return
				case addedStream <- i + additive:
					time.Sleep(1 * time.Second)
				}
			}
		}()
		return addedStream
	}

	done := make(chan interface{})
	intStream := generator(done, 1, 2, 3, 4)
	defer close(done)

	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	for v := range pipeline {
		fmt.Println(v)

	}

}
