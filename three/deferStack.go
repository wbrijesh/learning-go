package main

import "fmt"

func main() {
  fmt.Println("last in first out stack")

  for i := 0; i < 10; i++ {
    defer fmt.Println(i)
  }

  fmt.Println("This prints first")
}
