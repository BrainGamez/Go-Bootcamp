package main

import "fmt"

func main() {
	fmt.Println(f1())
}

func f1() int {
	return f2()
}

// de return krijgt een naam
func f2() (r int) {
	return 1
}

// doordat de ene functie de andere aanroept worden ze op de stack geplaatst
// wanneer het antwoord terug komt en de functie klaar is wordt deze weer van de stack gehaald
