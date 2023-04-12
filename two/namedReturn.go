package main

import "fmt"

func splitNumber(num int) (a float32, b float32) {
  a = (float32(num) / 2) + 1
  b = (float32(num) / 2) - 1
  return
}

func main() {
  fmt.Println(splitNumber(19))
}
