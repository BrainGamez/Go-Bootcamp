package main

import "fmt"

func main() {
	var yourName string
	fmt.Println("Enter you name: ")
	fmt.Scanf("%s", &yourName)
	fmt.Println("Is your name", yourName, "?")

}
