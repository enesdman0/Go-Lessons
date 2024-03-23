package main

import "fmt"

func main() {
	// underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println(underArray)
	// mySlc := underArray[2:5]
	// fmt.Println(mySlc)
	// mySlc2 := underArray[:6]
	// fmt.Println(mySlc2)
	// mySlc3 := underArray[6:]
	// fmt.Println(mySlc3)
	// mySlc4 := underArray[:]
	// fmt.Println(mySlc4)

	// fmt.Println(mySlc)
	// mySlc[0]=100
	// fmt.Println(mySlc)

	// mySlc := []int{1, 2, 3}
	// fmt.Println(mySlc)
	// mySlc = append(mySlc, 4, 5, 6, 7, 8, 9)
	// fmt.Println(mySlc)

	// mySlc2 := append(mySlc, 10, 11, 12)
	// fmt.Println(mySlc2)
	// mySlc[0]=100
	// fmt.Println(mySlc)
	// fmt.Println(mySlc2)

	// mySlc := []int{1, 2, 3}
	// mySlc2 := []int{4, 5}

	// mySlc = append(mySlc, mySlc2...)

	// fmt.Println(mySlc)

	mySlc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(mySlc)
	mySlc = mySlc[2:] // ilk 2 elemanı silme
	fmt.Println(mySlc)

	mySlc = mySlc[:len(mySlc)-2] // son 2 elemanı silme
	fmt.Println(mySlc)
}
