package main

import "fmt"

func main() {
	xs := map[[2]int]string{
		[2]int{1, 2}: "hello",
		[2]int{3, 4}: "world",
	}

	myMap, _ := xs[[2]int{1, 2}]

	fmt.Println(myMap)

	// fmt.Println(xs[[2]int{1, 2}], xs[[2]int{3, 4}])
	// fmt.Println(xs[[2]int{1}], xs[[2]int{4}])

}
