package main

import "fmt"

func main() {
	x := make(map[string]int)

	x["anything"] = 10
	x["something"] = 0

	// fmt.Println(x)
	var v = [3]int{}
	var ok = []bool{false, false, false}

	v[0], ok[0] = x["anything"]
	v[1], ok[1] = x["something"]
	v[2], ok[2] = x["something else"]

	// for i := 0; i < len(ok); i++ {
	// 	fmt.Println(ok[i])
	// }

	for i := range ok {
		fmt.Println(v[i], ok[i])
	}

	fmt.Println(x["something else"], x["something"], x["anything"])
}
