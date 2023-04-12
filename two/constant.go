package main 

import "fmt"

func main() {
  const pi float32 = 3.14

  // gives error
  // pi = 3.15

  fmt.Println("Value of pi is: ", pi )
}
