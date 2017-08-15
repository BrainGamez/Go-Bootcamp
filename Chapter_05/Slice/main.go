package main

import "fmt"

func main() {

	var x = []float64{1, 2, 3}
	x = make([]float64, 10)
	x[3] = 3
	x[6] = 1
	y := x[1:]
	fmt.Println(x, y)
}
