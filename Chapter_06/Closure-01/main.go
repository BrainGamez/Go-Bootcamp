package main

import "fmt"

func main() {
	funk := func(x, y int) int {
		return x + y
	}

	fmt.Println(funk(1, 2))
}
