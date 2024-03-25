// package main

// import "fmt"

// func main() {
// 	var employee struct {
// 		name      string
// 		age       int
// 		isMarried bool
// 	}
// 	fmt.Println(employee)

// 	employee.name="Enes"
// 	employee.age=20
// 	employee.isMarried=false

// 	fmt.Println(employee)
// }

package main

import "fmt"

type employee struct {
	name      string
	age       int
	isMarried bool
	kids      []string
}

func main() {
	// var e1 employee
	// e1.name = "Enes"
	// e1.age = 20
	// e1.isMarried = false

	// fmt.Println(e1)

	// var e2 employee
	// e2.name = "Buse"
	// e2.age = 20
	// e2.isMarried = false
	// fmt.Println(e2)

	e1 := employee{
		name:      "Enes",
		age:       20,
		isMarried: false,
	}
	fmt.Println(e1)
}
