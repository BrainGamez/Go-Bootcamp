package main

import "fmt"

var x = [4][4]int{
	{1, 2, 3, 4},
	{5, 6, 7, 8},
	{9, 10, 11, 12},
	{13, 14, 15, 16},
}

func main() {
	fmt.Println(x)
	fmt.Println(x[1][2])

	x[0][3] = 100

	fmt.Println(x)
}
