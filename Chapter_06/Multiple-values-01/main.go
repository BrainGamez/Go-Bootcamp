package main

import "fmt"

func f() (int, int) {
	return 5, 6
}

func main() {
	// opvallend is dat de x en y geinitialiseerd worden met de output van f()
	x, y := f()

	fmt.Println(x, y)
}
