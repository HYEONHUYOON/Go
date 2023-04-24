package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func main() {
	rand.Seed(time.Now().UnixNano())

	randNum := rand.Intn(100)
	count := 1

	for {
		fmt.Println("숫자를 입력하시오 : ")
		n, err := InputVal()
		if err != nil {
			fmt.Println("숫자만 입력하시오.")
		} else {
			if n > randNum {
				fmt.Println("다운")
			} else if n < randNum {
				fmt.Println("업")
			} else {
				fmt.Println("정답!! 시동한 수: ", count)
				break
			}
			count++
		}

	}
}

func InputVal() (int, error) {
	var n int
	//parsing
	//err => cant parsing to int64
	_, err := fmt.Scanln(&n)

	//버퍼를 비움
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}
