package main

import "fmt"

var p2kg = 0.45359237

func main() {
	var pound float64
	fmt.Print("Weight in pounds (lbs): ")
	_, err := fmt.Scanf("%f", &pound)
	if err != nil {
		fmt.Println(err)
	}
	kilo := pound * p2kg
	fmt.Println(pound, "(lbs) is equal to", kilo, "Kg")
	// *1.8 + 32
}
