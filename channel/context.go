package main

import (
	"context"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())

	// ctx, cancel = context.WithCancel(context.Background())
	// ctx = context.WithValue(ctx,"number",9)
	// ctx = context.WithValue(ctx,"number","Lilly")

	go Print(ctx)
	time.Sleep(5 * time.Second)

	// ctx.Done()
	cancel()

	wg.Wait()
}

func Print(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			println("tick")
		}
	}
}
