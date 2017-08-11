package main

import "fmt"

// 1 ft = 0.3048 m
var f2m = 0.3048

func main() {
	var foot float64
	fmt.Print("Distance in foot: ")
	_, err := fmt.Scanf("%f", &foot)
	if err != nil {
		fmt.Println(err)
	}
	meter := foot * f2m
	fmt.Println(foot, "foot is equal to", meter, "meter(s)")
	// *1.8 + 32
}
