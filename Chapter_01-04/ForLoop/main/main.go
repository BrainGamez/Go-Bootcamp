package main

import "fmt"

var waarde = "hoi"

func main() {
	var gek = true
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for !gek {
		verander()
		fmt.Println(waarde)
	}
}

func verander() {
	waarde = "doei"
}
