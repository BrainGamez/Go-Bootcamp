package main

import "fmt"

// de functie average heeft de 2 parameters: name a en type []float64
// het retourneert 1 waarde van float64 (aangegeven door de return)
func average(a []float64) float64 {
	// panic("not implemented")
	total := 0.0
	for _, v := range a {
		total += v
	}
	return total / float64(len(a))
}

// de functie main is het begin van het programma
func main() {
	// xs wordt geinitieerd met een slice van 5 float64's
	xs := []float64{98, 93, 77, 82, 83}
	// in de prnt wordt de functie average aangeroepen en de geretourneerde uitkomst wordt naar get scherm geprint
	fmt.Println(average(xs))
}
