package main

import "fmt"

func increaseNumbers() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := increaseNumbers()
	fmt.Println(increment())
	fmt.Println(increment())
}
