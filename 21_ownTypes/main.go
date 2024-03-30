package main

import "fmt"

type mile float64
type kilometer float64

// type myString string

func main() {
	// var m1 mile
	// m1 = 3.2
	// fmt.Println(m1)
	// fmt.Printf("%T %v\n", m1, m1)

	// f1 := float64(4.4)
	// fmt.Println(m1 + f1)

	// fmt.Printf("%T %v\n", m1+mile(f1), m1+mile(f1))
	// k1 := kilometer(7.8)
	// fmt.Printf("%T %v\n", k1, 11)
	// fmt.Println(m1 + k1)
	// m2 := mile(4.6)
	// fmt.Println(m1 * m2)
	// fmt.Printf("%T %v\n", m1+mile(m2), m1+mile(m2))
	// myString := "Enes"
	// fmt.Println(strings.ToUpper(myString))
	
	
	m1 := mile(10)
	k1 := kilometer(16)
	fmt.Println(toKilometer(m1))
	fmt.Println(toMile(k1))
}
func toKilometer(m mile) kilometer {
	return kilometer(m * 1.6)
}
func toMile(k kilometer) mile {
	return mile(k * 0.6)
}
