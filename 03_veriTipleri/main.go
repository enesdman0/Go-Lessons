package main

import "fmt"

func main() {
	name := "Enes"
	var age int16 = -256
	isMarried := false
	var weight float32 = 75.00

	isMarried = true

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMarried)
	fmt.Printf("%T\n", weight)
}
