package main

import (
	"Go/usepkg/custompkg"
	"Go/usepkg/exinit"
	"fmt"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	exinit.PrintD()

	data := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)

}
