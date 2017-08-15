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

func f(z int) {
	fmt.Println(z)
}

func main() {
	x := 5
	f(x)
}
