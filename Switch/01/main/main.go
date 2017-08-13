package main

import "fmt"

func main() {
	myVar := "test"

	switch myVar {
	case "nee":
		fmt.Println("nee klopt")
	case "ja":
		fmt.Println("ja klopt")
	default:
		println("klopt sowieso")
	}
}
