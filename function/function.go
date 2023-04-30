package main

import (
	"fmt"
	"os"
)

// 가변인자 함수1
func sum(nums ...int) int {
	sum := 0

	fmt.Printf("nums 타입 : %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

// 가변인자 함수 2 여러 타입 받기
func variable(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case bool:
			val := arg.(bool)
			fmt.Println(val)
		case float64:
			val := arg.(bool)
			fmt.Println(val)
		case int:
			val := arg.(bool)
			fmt.Println(val)
		case string:
			val := arg.(bool)
			fmt.Println(val)
		}
	}
}

func deferer(a int) {
	defer fmt.Println("defer")
	fmt.Println(a)
}

// 값이 캡쳐된다 결과 => 333
func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("Value Loop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

// 결과 => 012
func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("Value Loop2")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

type Writer func(string)

func WriteHello(writer Writer) {
	writer("Hello")
}

func main() {
	fmt.Println(sum(1, 23, 4))
	deferer(1)

	//람다
	f := func(a, b int) int {
		return a + b
	}

	fmt.Println(f(1, 2))

	CaptureLoop()
	CaptureLoop2()

	f, err := os.Create("test.tsxt")
	if err != nil {
		fmt.Println("file to create err")
		return
	}
	defer f.Close()

	WriteHello(func(msg string) {
		fmt.FPrintln(f, msg)
	})
}
