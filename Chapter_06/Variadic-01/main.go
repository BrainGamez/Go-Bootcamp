package main

import "fmt"

// ... betekend dat er een onbekend aantal van hetzelfde type geaccepteerd worden.
func average(a ...float64) float64 {
	total := 0.0
	for _, v := range a {
		total += v
	}
	return total / float64(len(a))
}

func main() {
	// xs := []float64{98, 93, 77, 82, 83}
	n := average(98, 93, 77, 82, 83)
	fmt.Println(n)
}
