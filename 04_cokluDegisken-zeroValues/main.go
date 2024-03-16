package main

import "fmt"

func main() {
	// var (
	// 	name      string = "Enes"
	// 	age              = 21
	// 	isMarried        = false

	// 	weight = 70
	// 	height = 176
	// )

	// var name, age, isMarried, weight, height = "Enes", 21, false, 70, 176
	// name, age, isMarried, weight, height := "Enes", 21, false, 70, 176

	var name string    // => string -> "",
	var age int        // => int -> 0,
	var isMarried bool // => bool -> false,
	var weight float32 // => float32 -> 0,

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)
	// fmt.Println(height)
}
