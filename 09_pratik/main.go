package main

import (
	"fmt"
	"strconv"
)

func main() {
	// -----------------------
	// 1 -) Type Conversion
	// x := 75
	// var y float64
	// y = float64(x)

	// fmt.Println(y)
	// -----------------------
	// 2 -) x, y= y, x
	// x := 5
	// y := 10
	// fmt.Print("x:", x, "\n", "y:", y, "\n")
	// x, y = y, x
	// fmt.Print("x:", x, "\n", "y:", y, "\n")
	// -----------------------
	// 3 -) None english variable name
	// yaş := 20
	// fmt.Println(yaş)
	// -----------------------
	// 4 -) Shadowing
	// x := 5
	// if true {
	// 	// x := 10
	// 	x = 10
	// 	x++
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)
	// -----------------------
	// 5 -) 40 as a string
	x := 65
	s := string(x)
	fmt.Printf("%T %v\n", x, x) // int 65
	fmt.Printf("%T %v\n", s, s) // string A

	y := strconv.Itoa(x)
	fmt.Printf("%T %v\n", y, y) // string "65"
}
