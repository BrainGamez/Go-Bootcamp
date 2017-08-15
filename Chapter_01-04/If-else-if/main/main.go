package main

import "fmt"

var myString = "eenString"

func main() {
	if myString == "5" {
		fmt.Println(myString)
	} else if myString == "eenString" {
		fmt.Println("poep " + myString)
	} else {
		fmt.Println("non Sense")
	}
}
