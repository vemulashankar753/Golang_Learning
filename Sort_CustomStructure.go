package main

import (
	"fmt"
	"sort"
)

type class struct {
	name         string
	employeelist []employee
}

type employee struct {
	name string
	id   int
}

// func (e *employee) sortBook(i, j int) bool {
// 	return e.books[i] < e.books[j]
// }

// type employeeList []employee

// func (e employeeList) Len() int {
// 	return len(e)
// }

// func (e employeeList) Less(i, j int) bool {
// 	return e[i].salary > e[j].salary
// }

// func (e employeeList) Swap(i, j int) {
// 	e[i], e[j] = e[j], e[i]
// }

func main() {

	c1 := class{
		name: "1stgrade",
		employeelist: []employee{
			{"shankar", 3},
			{"krishna", 1},
			{"rudra", 2},
			{"govind", 5},
		},
	}

	fmt.Printf("The value of the class c1: %+v \n ", c1)

	sort.Slice(c1.employeelist, func(i, j int) bool {
		return c1.employeelist[i].id > c1.employeelist[j].id
	})

	fmt.Printf("The value of the c1 after sorting employee list : %+v \n", c1)

}
