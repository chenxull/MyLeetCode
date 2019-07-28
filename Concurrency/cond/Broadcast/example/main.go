package main

import (
	"fmt"
	"sync"
)

type Button struct {
	//1
	Clicked *sync.Cond
}

func main() {

	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) { //2
		var tempwg sync.WaitGroup
		tempwg.Add(1)
		go func() {
			tempwg.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		tempwg.Wait()
	}

	var wg sync.WaitGroup //3
	wg.Add(3)
	subscribe(button.Clicked, func() { //4
		fmt.Println("Maximizing window.")
		wg.Done()
	})
	subscribe(button.Clicked, func() { //5
		fmt.Println("Displaying annoying dialog box!")
		wg.Done()
	})
	subscribe(button.Clicked, func() { //6
		fmt.Println("Mouse clicked.")
		wg.Done()
	})

	button.Clicked.Broadcast() //7

	wg.Wait()
}
