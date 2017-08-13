package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	miToKm  = 1.609344
	pToKg   = 0.45359237
	dType   = "<!DOCTYPE html>"
	htOpen  = "<html>"
	htClose = "</html>"
	hdOpen  = "<head>"
	hdClose = "</head>"
	byOpen  = "<body>"
	byClose = "</body>"
)

func main() {
	// fmt.Println("Choose an option:")
	// fmt.Println("1) Miles to Kilometers ")
	// fmt.Println("2) Pounds to Kilograms")
	// fmt.Println("3) Farenheit to Celsius")
	// var (
	// 	option int
	// 	// number float64
	// )

	// fmt.Scanln(&number)
	option, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Print("Enter the value: ")

	switch option {
	case 1:
		// fmt.Scanln(&number)
		number, err := strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(dType)
		fmt.Printf("%v\n", htOpen)
		fmt.Printf("\t%v%v\n", hdOpen, hdClose)
		fmt.Printf("\t%v\n", byOpen)
		fmt.Printf("\t\tMiles: %.0f<br>\n", number)
		fmt.Printf("\t\tKilometers: %.2f \n", number*miToKm)
		fmt.Printf("\t%v\n", byClose)
		fmt.Printf("%v\n", htClose)
	case 2:
		// fmt.Scanln(&number)
		number, err := strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(dType)
		fmt.Printf("%v\n", htOpen)
		fmt.Printf("\t%v%v\n", hdOpen, hdClose)
		fmt.Printf("\t%v\n", byOpen)
		fmt.Printf("\t\tPounds: %.0f<br>\n", number)
		fmt.Printf("\t\tKilograms: %.2f \n", number*pToKg)
		fmt.Printf("\t%v\n", byClose)
		fmt.Printf("%v\n", htClose)
	case 3:
		// fmt.Scanln(&number)
		number, err := strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(dType)
		fmt.Printf("%v\n", htOpen)
		fmt.Printf("\t%v%v\n", hdOpen, hdClose)
		fmt.Printf("\t%v\n", byOpen)
		fmt.Printf("\t\tFarenheit: %.0f<br>\n", number)
		fmt.Printf("\t\tCelsius: %.2f \n", (number-32)*5/9)
		fmt.Printf("\t%v\n", byClose)
		fmt.Printf("%v\n", htClose)
	default:
		fmt.Println("!!!", option, "is not an option!!!")
	}

}

// <!DOCTYPE html>
// 	<html>
// 		<head></head>
// 		<body>
// 			Miles: 50<br>
// 			Kilometers: 80.47
// 		</body>
// 	</html>
