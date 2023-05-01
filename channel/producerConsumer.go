package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var starTime = time.Now()

func main() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Println("Start Factory")

	wg.Add(3)
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go Paintcar(paintCh)

	wg.Wait()
	fmt.Println("End Factory")
}

func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	afer := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			car := &Car{}
			car.Body = "Sports"
			tireCh <- car
		case <-afer:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func InstallTire(tireCh, paintCh chan *Car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "Winter tire"
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

func Paintcar(paintCh chan *Car) {
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "Red"
		duraiton := time.Now().Sub(starTime)
		fmt.Printf("%.2f Complete Car : %s %s %s \n", duraiton.Seconds(), car.Body, car.Tire, car.Color)
	}
	wg.Done()
}
