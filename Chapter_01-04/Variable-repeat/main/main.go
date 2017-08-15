package main

import "fmt"

var myString string = "hallo"
var myInt int8 = 5
var myFloat float64 = 9.8

var (
	x = 8
	y = "dag"
)

func main() {
	i := "i"

	myBool, myInt32 := true, 5

	myString = "hoi"
	myInt = 40
	myInt32 = 2
	myFloat = 5.0
	myBool = false

	x = len("sjamalamadingdong")
	y, i = "hoi", "sniffels"

	fmt.Println(myString, myInt, myFloat, x, y, i, myBool, myInt32)

}
