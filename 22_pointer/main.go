package main

import "fmt"

func main() {
	// name := "enes"
	// fmt.Println(&name)
	// x := 22
	// fmt.Println(&x)
	// fmt.Println()
	// fmt.Printf("%T %v\n", x, x)
	// fmt.Printf("%T %v\n", &x, &x)
	// double(x)
	// fmt.Println(x)
	mySlc := []int{1, 10, 100}
	fmt.Println(mySlc)
	double(mySlc)
}
func double(slc []int) {
	for i := 0; i < len(slc); i++ {
		slc[i] *= 2
	}
	fmt.Println(slc)
}

// func double(x int) {
// 	x *= 2
// 	fmt.Println(x)
// }
