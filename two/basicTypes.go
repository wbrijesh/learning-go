package main

import "fmt"

var a bool = false
var b string = "hello"

// 0 to 255
var c uint8 = 120

// 0 to 65535
var d uint16 = 32000

// 0 to 4294967295
var e uint32 = 3294967295

// 0 to 18446744073709551615
var e uint64 = 18

// -128 to 127
var f int8 = -79

// -32768 to 32767
var g int16 = 100

// -2147483648 to 2147483647
var h int32 = 100

// -9223372036854775808 to 9223372036854775807
var i int64 = 100

var j float32 = 1.2

var k float64 = 1.2

var l complex64

var m complex128

// same as uint8
var n byte = 6 

// same as int32
var o rune = 3 


// int is either 32 or 64 bits
// uint is either 32 or 64 bits
// uintptr is unsigned integer to store value of pointer
