package main

import "fmt"

func main() {
  defer fmt.Println("Print me next.")

  fmt.Println("Print me first")
}
