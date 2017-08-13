package main

import "fmt"

func main() {
	myVar := "test"

	switch {
	case myVar == "nee":
		fmt.Println("nee klopt")
	case myVar == "ja":
		fmt.Println("ja klopt")
	default:
		println("klopt sowieso")
	}
}
