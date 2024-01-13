package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type Person struct {
	name string
	age int
}

type Class struct {
	persons []Person
	classNo  int
}
type College struct{
	classes []Class
	clgName string
}

func main() {

	fmt.Println("The Program usage")
	p1 := Person{name: "shan", age: 15}
	p2 := Person{name: "krish", age: 17}

    c1 := Class{persons: []Person{p1,p2}, classNo : 1}

	clg := College{classes: []Class{c1}, clgName: "JNTU"}

    spew.Sdump(clg)

}

