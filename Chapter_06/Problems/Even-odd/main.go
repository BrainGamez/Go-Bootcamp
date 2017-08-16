package main

import "fmt"

// Write a function which takes an integer and halves it and returns true if it was even or false if it was odd. For example half(1) should return (0, false) and half(2) should return (1, true).

func funk(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func main() {
	var x int
	fmt.Printf("voer een getal in: ")
	fmt.Scanln(&x)
	fmt.Println(funk(x))
}
