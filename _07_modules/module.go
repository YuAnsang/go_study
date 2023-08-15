package main

import (
	"fmt"
	"modules/_07_modules/custompkg"

	"github.com/guptarohit/asciigraph"
)

func main() {
	fmt.Println("테스트")
	custompkg.PrintCustom()

	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
