package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Millisecond) {

			cadence.Broadcast()
		}
	}()

	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}

	/*
		1. tryDir 允许一个人尝试向某个方向移动并返回，无论他们是否成功。 每个方向都表示为试图朝这个方向移动的次数。
		2. 首先，我们通过将该方向递增1来朝着某个方向移动。 我们将在第3章详细讨论atomic包。现在，你只需要知道这个包的操作是原子操作。
		3. 每个人必须以相同的速度或节奏移动。 takeStep模拟所有动作之间的恒定节奏。
		4. 在这里，这个人意识到他们不能在这个方向上放弃。 我们通过将该方向递减1来表示这一点。
	*/
	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool { //1
		fmt.Fprintf(out, " %v", dirName)
		atomic.AddInt32(dir, 1) //2
		takeStep()              //3
		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprint(out, ". Success!")
			return true
		}
		takeStep()
		atomic.AddInt32(dir, -1) //4
		return false
	}

	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool { return tryDir("left", &left, out) }
	tryRight := func(out *bytes.Buffer) bool { return tryDir("right", &right, out) }

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() { fmt.Println(out.String()) }()
		defer walking.Done()
		fmt.Fprintf(&out, "%v is trying to scoot:", name)
		for i := 0; i < 5; i++ { //1
			if tryLeft(&out) || tryRight(&out) { //2
				return
			}
		}
		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
	}

	var peopleInHallway sync.WaitGroup //3
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()

}
