package main

import "fmt"

// func sum(input ...int) int {
func sum(input []int) int {
	total := 0

	for v := range input {
		total += v
	}
	return total
}

func main() {
	// x := sum(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	x := sum([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	fmt.Println(x)
}
