package main

import "fmt"

func main() {
	//Long declaration
	var i int
	var b bool
	var str string
	var c float64

	//Shor declaration
	//i := 1
	//b := true
	//str : ="shankar"

	fmt.Printf("Var a %T = %d", i, i)
	fmt.Printf("Var a %T = %q", str, str)
	fmt.Printf("Var a %T = %f", c, c)
	fmt.Printf("Var a %T = %t", b, b)

}

//There are only 25 keywords in Go. These keywords may not be used as identifiers.

//break        default      func         interface    select
//case         defer        go           map          struct
//chan         else         goto         package      switch
//const        fallthrough  if           range        type
//continue     for          import       return       var
