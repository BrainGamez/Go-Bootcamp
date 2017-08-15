package main

import "fmt"

var (
	miToKm  = 1.609344
	pToKg   = 0.45359237
	devider = "+-------------------------+"
)

func main() {
	fmt.Println("Choose an option:")
	fmt.Println("1) Miles to Kilometers ")
	fmt.Println("2) Pounds to Kilograms")
	fmt.Println("3) Farenheit to Celsius")
	var (
		option int
		number float64
	)

	fmt.Scanln(&option)

	switch option {
	case 1:
		fmt.Println(devider)
		fmt.Print("  Miles: ")
		fmt.Scanln(&number)
		fmt.Println(devider)
		fmt.Printf("  Kilometers: %.2f \n", number*miToKm)
		fmt.Println(devider)
	case 2:
		fmt.Println(devider)
		fmt.Print("  Pounds: ")
		fmt.Scanln(&number)
		fmt.Println(devider)
		fmt.Printf("  Kilograms: %.2f \n", number*pToKg)
		fmt.Println(devider)
	case 3:
		fmt.Println(devider)
		fmt.Print("  Pounds: ")
		fmt.Scanln(&number)
		fmt.Println(devider)
		fmt.Printf("  Kilograms: %.2f \n", (number-32)*5/9)
		fmt.Println(devider)
	default:
		fmt.Println("!!!", option, "is not an option!!!")
	}

}
