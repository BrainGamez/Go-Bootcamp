package main

import "fmt"

func devide() func(f int) string {
	x := 10

	return func(a int) string {

		var b string
		if a%x == 0 {
			b = "deelbaar door 10"
		} else {
			b = "wtf!"
		}

		return b
	}
}

func main() {
	// uitkomst wordt geinitialiseerd met de functie welke de return is van devide
	// uitkomst := func(f int) string {........
	uitkomst := devide()
	fmt.Printf("%T\n", uitkomst)
	fmt.Println(uitkomst(99))
}
