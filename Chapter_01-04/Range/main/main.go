package main

import "fmt"

func main() {
	for range "weergave" {
		fmt.Println("klopt")

	}
	myRange()
}

func myRange() {
	for i, c := range "weergave en 汉语 !" {
		fmt.Printf("%v %c\n", i, c)
	}
	for _, c := range "بسم الله الرحمن الرحيم" {
		fmt.Printf("%c\n", c)
	}
}
