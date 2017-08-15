package main

import "fmt"

// 1 ft = 0.3048 m
var m2k = 1.609344

func main() {
	var mile float64
	fmt.Print("Distance in foot: ")
	_, err := fmt.Scanf("%f", &mile)
	if err != nil {
		fmt.Println(err)
	}
	kmeter := mile * m2k
	fmt.Println("+-------------------------+")
	fmt.Printf("| Miles: %.0f               |\n", mile)
	fmt.Printf("| Kilometers: %.2f       |\n", kmeter)
	fmt.Println("+-------------------------+")
	// *1.8 + 32
}
