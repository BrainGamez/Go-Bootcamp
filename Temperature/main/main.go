package main

import "fmt"

func main() {
	var celsius float64
	fmt.Print("Temp in °C: ")
	fmt.Scanf("%f", &celsius)
	farenheit := celsius*1.8 + 32
	fmt.Println("Temp in Farenheit:", farenheit, "°F")
	// *1.8 + 32
}
