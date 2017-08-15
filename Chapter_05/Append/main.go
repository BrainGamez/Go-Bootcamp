package main

import "fmt"

func main() {
	x := []float64{1, 2, 3, 4, 5}

	fmt.Println(x)

	x = append(x, 6)

	fmt.Println(x)

	y := make([]float64, 3)
	y[0] = 8
	y[1] = 9

	copy(y[2:], x)

	fmt.Println(y)

	x = append(x, y[1:]...)

	fmt.Println(x)

	// z := make([]float64, 10)
	z := []float64{1}

	copy(z, x)

	fmt.Println(z)
}
