package main

import (
	"fmt"
	"sync"
	"time"
)

func squre(wg *sync.WaitGroup, c chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-terminate:
			fmt.Println("Terminating")
			wg.Done()
			return
		case n := <-c:
			fmt.Println("Squre", n*n)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go squre(&wg, ch)

	ch <- 9
	wg.Wait()

}
