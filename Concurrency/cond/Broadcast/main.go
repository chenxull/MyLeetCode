package main

import (
	"fmt"
	"sync"
)

type Button struct {
	Clicked *sync.Cond
}

func main() {
	button := Button{
		Clicked: sync.NewCond(&sync.Mutex{}),
	}

	subscribe := func(c *sync.Cond, fn func()) {
		var temwg sync.WaitGroup
		temwg.Add(1)
		go func() {
			temwg.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		temwg.Wait()
	}

	var wg sync.WaitGroup
	wg.Add(3)

	subscribe(button.Clicked, func() {
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

	button.Clicked.Broadcast()
	wg.Wait()

}
