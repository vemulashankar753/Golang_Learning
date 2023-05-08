package main

import "fmt"

// section: var
var s string    // defaults to ""
var r rune      // defaults to 0
var bt byte     // defaults to 0
var i int       // defaults to 0
var ui uint     // defaults to 0
var f float32   // defaults to 0
var c complex64 // defaults to 0
var b bool      // defaults to false
var arr [2]int  // defaults to [0 0]
var obj struct {
	b   bool
	arr [2]int
}                     // defaults to {false [0 0]}
var si []int          // defaults to int
var ch chan string    // defaults to nil
var mp map[int]string // defaults to nil
var fn func()         // defaults to nil
var ptr *string       // defaults to nil
// section: var

func main() {
	fmt.Println("string              :", s)
	fmt.Println("rune                :", r)
	fmt.Println("byte                :", bt)
	fmt.Println("int                 :", i)
	fmt.Println("uint                :", ui)
	fmt.Println("float32             :", f)
	fmt.Println("complex64           :", c)
	fmt.Println("bool                :", b)
	fmt.Println("array [2]int        :", arr)
	fmt.Println("struct              :", obj)
	fmt.Println("slice []int         :", si)
	fmt.Println("channel chan string :", ch)
	fmt.Println("map map[int]string  :", mp)
	fmt.Println("function func()     :", fn)
	fmt.Println("*string             :", ptr)

}
