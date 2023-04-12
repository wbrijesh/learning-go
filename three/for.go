package main 

import "fmt"
// import "time"

func main() {
  for i := 0; i < 1000000; i++ {
    fmt.Println("I is: ", i)
  }

  j := 0
  for j < 10000 {
    j++
    fmt.Println("J is: ", j)
  }
}
