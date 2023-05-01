package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex sync.Mutex

var wg sync.WaitGroup

func Printhangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
	wg.Done()
}

func PrintNum() {
	for i := 0; i <= 10; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func main() {
	wg.Add(2)

	go PrintNum()
	go Printhangul()

	wg.Wait()

	fmt.Println()

	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go dinungProblem("A", fork, spoon, "포크", "수저")
	go dinungProblem("B", spoon, fork, "수저", "포크")

	wg.Wait()
}

func dinungProblem(name string, first, second *sync.Mutex, firstname, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다.\n", name)
		first.Lock()
		fmt.Printf("%s %s 획득.\n", name, firstname)
		first.Lock()
		fmt.Printf("%s %s 획득.\n", name, secondName)

		fmt.Printf("%s 밥을 먹습니다.\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}
