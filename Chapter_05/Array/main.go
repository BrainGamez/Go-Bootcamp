package main

import "fmt"

func main() {

  var x = [50]float64 {
    98,
    93,
    77,
    82,
    83,
  }

  var total float64 = 0

  for _, v := range x {
    total += v
  }

  // for i := 0; i < len(x); i++ {
  //   total += x[i]
  // }

  fmt.Println(total / float64(len(x)))

}
