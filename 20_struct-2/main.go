package main

import "fmt"

type employee struct {
	name      string
	age       int
	isMarried bool
}
type manager struct {
	employee
	hasDegree bool
}

func main() {
	// m1 := manager{
	// 	employee: employee{
	// 		name:      "Buse",
	// 		age:       20,
	// 		isMarried: false,
	// 	},
	// 	hasDegree: true,
	// }
	// fmt.Println(m1)

	// m1 := manager{}
	// m1.name = "Buse"
	// m1.age = 20
	// m1.isMarried = false
	// m1.hasDegree = true
	// fmt.Println(m1)
	
	// Anonim Struct
	theBoss := struct {
		name  string
		money bool
		}{name: "THE BOSS", money: true}
		fmt.Println(theBoss)
}
