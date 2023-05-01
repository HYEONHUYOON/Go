package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

const M = 10

func hash(d int) int {
	return d % M
}

func main() {
	m := make(map[int]Product)

	m[1] = Product{Name: "Apple", Price: 100}
	m[2] = Product{Name: "DDD", Price: 200}
	m[3] = Product{Name: "Grapefruit", Price: 1120}
	m[4] = Product{Name: "Orange", Price: 140}

	//순서 보장이 안됌
	for key, value := range m {
		fmt.Println(key, value)
	}

	delete(m, 1)
	value, ok := m[1]

	fmt.Println(value, ok)

	h := [M]string{}
	h[hash(23)] = "가"
	h[hash(259)] = "나"
}
