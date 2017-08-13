package main

import (
	"fmt"
	"strconv"
)

// func main() {
// 	fmt.Println(strconv.ParseInt("0D0A", 16, 64))
// 	fmt.Println(strconv.FormatInt(3338, 16))
// 	fmt.Println(strconv.ParseInt(strconv.FormatInt(3338, 16), 16, 64))
// }

func main() {
	fmt.Println(
		strconv.ParseInt(
			strconv.FormatInt(3338, 16),
			16,
			64,
		),
	)
}
