package main

import "fmt"

func main() {

	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	// name, _ := elements["N"]
	// fmt.Println(name)

	fmt.Printf("Pick a known element: ")
	var el string
	fmt.Scanln(&el)
	name, _ := elements[el]
	fmt.Printf("You picked %v\n", name)

}
