package main

import "fmt"

func main() {
	x, y := 15, 10

	fmt.Printf("x: %T, %v\n", x, x)
	fmt.Printf("y: %T, %v\n", y, y)

	fmt.Printf("+: %T, %v\n", (x + y), (x + y))
	fmt.Printf("-: %T, %v\n", (x - y), (x - y))
	fmt.Printf("/: %T, %v\n", (x / y), (x / y))
	fmt.Printf("/: %T, %v\n", (x % y), (x % y))

	// z := float64(5 / 2)
	z := float64(5 / 2.0)
	fmt.Printf("%T %v", z, z)
}
