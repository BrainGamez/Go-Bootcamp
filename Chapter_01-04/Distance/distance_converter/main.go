package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Create a program which parses a query to do distance conversions. For example, from a terminal:
// 	$ distance_converter 50mi km
// Should produce:
// 	80.47km
// It should support miles (mi), kilometers (km), feet (ft) and meters (m).
const (
	miToKm = 1.609344
	miToFt = 5280
	miToMe = 1609.344

	kmToMi = 0.621371192
	kmToFt = 3280.8399
	kmToMe = 1000

	ftToMi = 0.000189393939
	ftToKm = 0.0003048
	ftToMe = 0.3048

	meToMi = 0.000621371192
	meToKm = 0.001
	meToFt = 3.2808399
)

var from string
var to string
var input float64

func main() {

	from := os.Args[1]

	to := os.Args[2]

	// input, _ := strconv.ParseFloat(from[0:len(from)-2], 64)

	if strings.HasSuffix(from, "mi") {
		// input, _ := strconv.ParseFloat(strings.TrimRight(os.Args[1], "mi"), 64)
		input, _ := strconv.ParseFloat(from[0:len(from)-2], 64)
		switch to {
		case "km":
			fmt.Printf("%.2f miles is equal to %.2f kilometers\n", input, input*miToKm)
		case "ft":
			fmt.Printf("%.2f miles is equal to %.2f feet\n", input, input*miToFt)
		case "m":
			fmt.Printf("%.2f miles is equal to %.2f meters\n", input, input*miToMe)
		default:
			fmt.Printf("%.2f miles\n", input)
		}
	} else if strings.HasSuffix(from, "km") {
		// input, _ := strconv.ParseFloat(strings.TrimRight(os.Args[1], "km"), 64)
		input, _ := strconv.ParseFloat(from[0:len(from)-2], 64)
		switch to {
		case "mi":
			fmt.Printf("%.2f kilometers is equal to %.2f miles\n", input, input*kmToMi)
		case "ft":
			fmt.Printf("%.2f kilometers is equal to %.2f feet\n", input, input*kmToFt)
		case "m":
			fmt.Printf("%.2f kilometers is equal to %.2f meters\n", input, input*kmToMe)
		default:
			fmt.Printf("%.2f kilometers\n", input)
		}
	} else if strings.HasSuffix(from, "ft") {
		// input, _ := strconv.ParseFloat(strings.TrimRight(os.Args[1], "ft"), 64)
		input, _ := strconv.ParseFloat(from[0:len(from)-2], 64)
		switch to {
		case "mi":
			fmt.Printf("%.2f feet is equal to %.2f miles\n", input, input*ftToMi)
		case "km":
			fmt.Printf("%.2f feet is equal to %.2f kilometers\n", input, input*ftToKm)
		case "m":
			fmt.Printf("%.2f feet is equal to %.2f meters\n", input, input*ftToMe)
		default:
			fmt.Printf("%.2f feet\n", input)
		}
	} else if strings.HasSuffix(from, "m") {
		// input, _ := strconv.ParseFloat(strings.TrimRight(os.Args[1], "m"), 64)
		input, _ := strconv.ParseFloat(from[0:len(from)-1], 64)
		switch to {
		case "mi":
			fmt.Printf("%.2f meters is equal to %.2f miles\n", input, input*meToMi)
		case "km":
			fmt.Printf("%.2f meters is equal to %.2f kilometers\n", input, input*meToKm)
		case "ft":
			fmt.Printf("%.2f meters is equal to %.2f feet\n", input, input*meToFt)
		default:
			fmt.Printf("%.2f meters\n", input)
		}
	} else {
		fmt.Println("Not a valid input!")
	}
}
