package main

import "fmt"

// dit werkt niet
// func f() {
// 	fmt.Println(x)
// }
//
// func main() {
// 	x := 5
// 	f()
// }

var x int

func f() {
	fmt.Println(x)
}

func main() {
	x = 5
	f()
}
